package main

import (
	"fmt"
	"gin-learn/gin-todolist/user/config"
	"gin-learn/gin-todolist/user/discovery"
	"gin-learn/gin-todolist/user/internal/handler"
	"gin-learn/gin-todolist/user/internal/repository"
	"gin-learn/gin-todolist/user/internal/service"
	"net"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	config.InitConfig()
	repository.InitDB()

	// 启动grpc服务
	grpcAddr := viper.GetString("server.grpcAddress")
	server := grpc.NewServer()
	defer server.Stop()
	service.RegisterUserServiceServer(server, handler.NewUserService())
	reflection.Register(server)
	listen, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		panic(err)
	}

	// etcd地址
	etcdAddrs := viper.GetStringSlice("etcd.address")
	// 创建注册器
	etcdRegister := discovery.NewRegister(etcdAddrs, logrus.New())
	defer etcdRegister.Stop()
	// 服务注册键值对
	userNode := discovery.Server{
		Name: viper.GetString("server.domain"),
		Addr: grpcAddr,
	}
	if _, err := etcdRegister.RegisterServer(userNode, 10); err != nil {
		panic(fmt.Sprintf("start server failed, err: %v", err))
	}

	// 开始监听
	logrus.Info("server started listen on ", grpcAddr)
	if err := server.Serve(listen); err != nil {
		panic(err)
	}
}
