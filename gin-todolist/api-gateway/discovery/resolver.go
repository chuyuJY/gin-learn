package discovery

import (
	"context"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc/resolver"
)

const (
	// etcd resolver 负责的 scheme 类型
	scheme = "etcd"
)

// gRPC 的 Resolver
type EtcdResolver struct {
	scheme      string
	EtcdAddrs   []string
	DialTimeout int

	closeCh      chan struct{}
	watchCh      clientv3.WatchChan
	client       *clientv3.Client
	keyPrefix    string
	srvAddrsList []resolver.Address

	cc     resolver.ClientConn
	logger *logrus.Logger
}

// 1. 实现 grpc.Resolver 接口
// interface
func (r *EtcdResolver) ResolveNow(o resolver.ResolveNowOptions) {

}

// interface
func (r *EtcdResolver) Close() {
	r.closeCh <- struct{}{}
}

// 2. 实现grpc.Builder接口
// interface
func (r *EtcdResolver) Scheme() string {
	return r.scheme
}

// interface
func (r *EtcdResolver) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	r.cc = cc

	r.keyPrefix = BuildPrefix(&Server{
		Name:    target.Endpoint,
		Version: target.Authority,
	})
	// fmt.Println("target.Endpoint: ", target.Endpoint)
	// fmt.Println("target.URL.Path: ", target.URL.Path)
	// fmt.Println("r.keyPrefix: ", r.keyPrefix)
	if err := r.start(); err != nil {
		return nil, err
	}
	return r, nil
}

func NewResolver(etcdAddrs []string, logger *logrus.Logger) *EtcdResolver {
	return &EtcdResolver{
		scheme:      scheme,
		EtcdAddrs:   etcdAddrs,
		DialTimeout: 3,
		logger:      logger,
	}
}

// etcd注册
func (r *EtcdResolver) start() error {
	var err error
	r.client, err = clientv3.New(clientv3.Config{
		Endpoints:   r.EtcdAddrs,
		DialTimeout: time.Duration(r.DialTimeout) * time.Second,
	})
	if err != nil {
		return err
	}

	resolver.Register(r)

	r.closeCh = make(chan struct{})

	// 启动时, 先同步一下
	if err = r.sync(); err != nil {
		return err
	}

	go r.watch()

	return nil
}

// watch update events
func (r *EtcdResolver) watch() {
	// 定时器, 周期为 1min, 此处用来定时同步
	ticker := time.NewTicker(time.Minute)

	// 设置监听 etcd 相应的 key prefix, 变更事件发生时, 更新本地缓存
	r.watchCh = r.client.Watch(context.Background(), r.keyPrefix, clientv3.WithPrefix())

	for {
		select {
		case <-r.closeCh:
			return
		case res, ok := <-r.watchCh:
			if ok {
				r.update(res.Events)
			}
		case <-ticker.C:
			if err := r.sync(); err != nil {
				r.logger.Error("sync failed", err)
			}
		}
	}
}

// update
func (r *EtcdResolver) update(events []*clientv3.Event) {
	for _, event := range events {
		var server *Server
		var err error

		switch event.Type {
		case clientv3.EventTypePut: // 新增本地地址
			server, err = ParseValue(event.Kv.Value)
			if err != nil {
				continue
			}
			addr := resolver.Address{
				Addr:     server.Addr,
				Metadata: server.Weight,
			}
			if !Exist(r.srvAddrsList, addr) {
				r.srvAddrsList = append(r.srvAddrsList, addr)
				r.cc.UpdateState(resolver.State{
					Addresses: r.srvAddrsList,
				})
			}
		case clientv3.EventTypeDelete: // 删除本地地址
			server, err = SplitPath(string(event.Kv.Key))
			if err != nil {
				continue
			}
			addr := resolver.Address{
				Addr: server.Addr,
			}
			if s, ok := Remove(r.srvAddrsList, addr); ok {
				r.srvAddrsList = s
				r.cc.UpdateState(resolver.State{
					Addresses: r.srvAddrsList,
				})
			}
		}
	}
}

// sync 同步获取所有地址信息
func (r *EtcdResolver) sync() error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	res, err := r.client.Get(ctx, r.keyPrefix, clientv3.WithPrefix())
	if err != nil {
		return err
	}

	r.srvAddrsList = []resolver.Address{}
	for _, v := range res.Kvs {
		server, err := ParseValue(v.Value)
		if err != nil {
			continue
		}
		fmt.Println("得到服务地址为: ", server.Addr)
		addr := resolver.Address{
			Addr:     server.Addr,
			Metadata: server.Weight,
		}
		r.srvAddrsList = append(r.srvAddrsList, addr)
	}
	r.cc.UpdateState(resolver.State{
		Addresses: r.srvAddrsList,
	})
	return nil
}
