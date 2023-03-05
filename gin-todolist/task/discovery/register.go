package discovery

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type EtcdRegister struct {
	cli         *clientv3.Client                        // etcd client
	leasesID    clientv3.LeaseID                        // 租约
	ctx         context.Context                         // 上下文
	cancel      context.CancelFunc                      // 取消
	closeCh     chan struct{}                           // 是否关闭
	keepAliveCh <-chan *clientv3.LeaseKeepAliveResponse // 心跳检测

	EtcdAddrs   []string       // etcd地址
	DialTimeout int            // 超时时间
	logger      *logrus.Logger // 日志
	srvInfo     Server         // 服务信息
	srvTTL      int64          // 租约时间
}

// 基于ETCD创建一个register
func NewRegister(etcdAddrs []string, logger *logrus.Logger) *EtcdRegister {
	return &EtcdRegister{
		EtcdAddrs:   etcdAddrs,
		DialTimeout: 3,
		logger:      logger,
	}
}

// 注册gRPC服务
func (r *EtcdRegister) RegisterServer(srvInfo Server, ttl int64) (chan<- struct{}, error) {
	if strings.Split(srvInfo.Addr, ":")[0] == "" {
		return nil, errors.New("invalid ip address")
	}

	var err error
	// 初始化客户端
	if r.cli, err = clientv3.New(clientv3.Config{
		Endpoints:   r.EtcdAddrs,
		DialTimeout: time.Duration(r.DialTimeout) * time.Second,
	}); err != nil {
		return nil, err
	}
	r.srvInfo = srvInfo
	r.srvTTL = ttl
	r.ctx, r.cancel = context.WithCancel(context.Background())

	if err = r.register(); err != nil {
		return nil, err
	}

	r.closeCh = make(chan struct{})
	go r.Watcher()
	return r.closeCh, nil
}

// 初始化ETCD自带的实例
func (r *EtcdRegister) register() error {
	// 创建租约, 大小为 r.srvTTL秒
	if err := r.CreateLease(); err != nil {
		return err
	}

	// 绑定租约
	if err := r.BindLease(); err != nil {
		return err
	}

	// 自动续约
	if err := r.KeepAlive(); err != nil {
		return err
	}

	return nil
}

// 创建租约
func (r *EtcdRegister) CreateLease() error {
	leaseResp, err := r.cli.Grant(r.ctx, r.srvTTL)
	if err != nil {
		log.Printf("createLease failed,error %v \n", err)
		return err
	}
	r.leasesID = leaseResp.ID
	return nil
}

// 绑定租约
// 将租约和对应的KEY-VALUE绑定
func (r *EtcdRegister) BindLease() error {
	data, err := json.Marshal(r.srvInfo)
	if err != nil {
		return err
	}
	res, err := r.cli.Put(r.ctx, BuildRegisterPath(&r.srvInfo), string(data), clientv3.WithLease(r.leasesID))
	if err != nil {
		log.Printf("bindLease failed,error %v \n", err)
		return err
	}
	log.Printf("bindLease success %v \n", res)
	return nil
}

// 续租 定时发送心跳, 表示服务存活
func (r *EtcdRegister) KeepAlive() error {
	var err error
	r.keepAliveCh, err = r.cli.KeepAlive(r.ctx, r.leasesID)
	return err
}

// 监测服务状况
func (r *EtcdRegister) Watcher() error {
	ticker := time.NewTicker(time.Duration(r.srvTTL) * time.Second)
	for {
		select {
		case <-r.closeCh:
			if err := r.UnRegisterServer(); err != nil {
				fmt.Println("unregister failed error: ", err)
			}
			if _, err := r.cli.Revoke(context.Background(), r.leasesID); err != nil {
				fmt.Println("revoke fail")
			}
		case res := <-r.keepAliveCh:
			if res == nil {
				if err := r.register(); err != nil {
					fmt.Println("register fail")
				}
			}
		case <-ticker.C:
			if r.keepAliveCh == nil {
				if err := r.register(); err != nil {
					fmt.Println("register fail")
				}
			}
		}
	}
}

func (r *EtcdRegister) Stop() {
	r.cancel()
	r.closeCh <- struct{}{}
}

func (r *EtcdRegister) UnRegisterServer() error {
	_, err := r.cli.Delete(context.Background(), BuildRegisterPath(&r.srvInfo))
	return err
}

func (r *EtcdRegister) UpdateHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		weightstr := req.URL.Query().Get("weight")
		weight, err := strconv.Atoi(weightstr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(err.Error()))
			return
		}

		var update = func() error {
			r.srvInfo.Weight = int64(weight)
			data, err := json.Marshal(r.srvInfo)
			if err != nil {
				return err
			}

			_, err = r.cli.Put(context.Background(), BuildRegisterPath(&r.srvInfo), string(data), clientv3.WithLease(r.leasesID))
			return err
		}

		if err := update(); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(err.Error()))
			return
		}

		_, _ = w.Write([]byte("update server weight success"))
	}
}

func (r *EtcdRegister) GetServerInfo() (Server, error) {
	resp, err := r.cli.Get(context.Background(), BuildRegisterPath(&r.srvInfo))
	if err != nil {
		return r.srvInfo, err
	}

	server := Server{}
	if resp.Count >= 1 {
		if err := json.Unmarshal(resp.Kvs[0].Value, &server); err != nil {
			return server, err
		}
	}

	return server, err
}
