package main

import (
	"fmt"
	"gin-learn/gin-todolist/task/config"
	"gin-learn/gin-todolist/task/discovery"
	"gin-learn/gin-todolist/task/internal/handler"
	"gin-learn/gin-todolist/task/internal/repository"
	"gin-learn/gin-todolist/task/internal/service"
	"net"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func main() {
	config.InitConfig()
	repository.InitDB()

	// 启动grpc服务
	grpcAddr := viper.GetString("server.grpcAddress")
	server := grpc.NewServer()
	defer server.Stop()
	service.RegisterTaskServiceServer(server, handler.NewTaskService())
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
	taskNode := discovery.Server{
		Name: viper.GetString("server.domain"),
		Addr: grpcAddr,
	}
	if _, err := etcdRegister.RegisterServer(taskNode, 10); err != nil {
		panic(fmt.Sprintf("start server failed, err: %v", err))
	}

	// 开始监听
	logrus.Info("server started listen on ", grpcAddr)
	if err := server.Serve(listen); err != nil {
		panic(err)
	}
}
