package main

import (
	"context"
	"fmt"
	"gin-learn/gin-todolist/api-gateway/config"
	"gin-learn/gin-todolist/api-gateway/discovery"
	"gin-learn/gin-todolist/api-gateway/internal/service"
	"gin-learn/gin-todolist/api-gateway/pkg/utils"
	"gin-learn/gin-todolist/api-gateway/routes"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
)

func main() {
	config.InitConfig()
	go startListen() // 转载路由
	{
		osSignals := make(chan os.Signal, 1)
		signal.Notify(osSignals, os.Interrupt, os.Kill, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
		s := <-osSignals
		fmt.Println("exit!", s)
	}
}

func startListen() {
	// etcd注册
	etcdAddrs := viper.GetStringSlice("etcd.address")
	etcdRegister := discovery.NewResolver(etcdAddrs, logrus.New())
	resolver.Register(etcdRegister)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// 服务名
	userServiceName := viper.GetString("domain.user")
	taskServiceName := viper.GetString("domain.task")

	// RPC连接, 还会建立服务发现
	connUser, err := RPCConnect(ctx, userServiceName, etcdRegister)
	if err != nil {
		return
	}
	userService := service.NewUserServiceClient(connUser)

	connTask, err := RPCConnect(ctx, taskServiceName, etcdRegister)
	if err != nil {
		return
	}
	taskService := service.NewTaskServiceClient(connTask)

	ginRouter := routes.NewRouter(userService, taskService)
	server := &http.Server{
		Addr:           viper.GetString("server.port"),
		Handler:        ginRouter,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err = server.ListenAndServe()
	if err != nil {
		fmt.Printf("绑定HTTP到 %s 失败！可能是端口已经被占用，或用户权限不足", server.Addr)
	}
	fmt.Println("gateway listen on:", server.Addr)
	go func() {
		// 优雅关闭
		utils.GracefullyShutdown(server)
	}()
}

func RPCConnect(ctx context.Context, serviceName string, etcdRegister *discovery.EtcdResolver) (conn *grpc.ClientConn, err error) {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	target := fmt.Sprintf("%s:///%s", etcdRegister.Scheme(), serviceName)
	// fmt.Println("addr: ", target)
	conn, err = grpc.DialContext(ctx, target, opts...) // 跳转到 Build Resplver
	if err != nil {
		log.Panic("grpc.Dial failed: ", err)
	}
	return
}
