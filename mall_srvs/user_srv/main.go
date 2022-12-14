package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"

	"mall_srvs/user_srv/global"
	"mall_srvs/user_srv/handler"
	"mall_srvs/user_srv/initialize"
	"mall_srvs/user_srv/proto"
	"mall_srvs/user_srv/srv_config"
)

func main() {
	// logger
	initialize.InitLogger()
	// init config
	//initialize.InitConfig()

	// init nacos config
	initialize.InitNacosConfig()
	// init config
	initialize.InitNacosToConfig()

	// mysql
	initialize.InitMySQL()

	server := grpc.NewServer()
	proto.RegisterUserServer(server, &handler.UserServer{})
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d",
		global.ServerConfig.Host, global.ServerConfig.Port))
	if err != nil {
		panic("failed to listen:" + err.Error())
	}
	zap.S().Infof("user_srv is running,the ip is %v port is %v",
		global.ServerConfig.Host, global.ServerConfig.Port)

	// 注册健康检查服务
	grpc_health_v1.RegisterHealthServer(server, health.NewServer())

	// 服务注册  srv to consul
	client, serviceID := srv_config.RegisterToConsul()

	go func() {
		err = server.Serve(lis)
		if err != nil {
			panic("failed to start user_srv." + err.Error())
		}
	}()

	// 优雅退出  关闭后丢弃consul配置信息
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	if err = client.Agent().ServiceDeregister(serviceID); err != nil {
		zap.S().Info("deregister consul failed")
	}
	zap.S().Info("deregister consul successful")
}
