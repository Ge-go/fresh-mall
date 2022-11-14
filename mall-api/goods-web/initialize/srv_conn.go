package initialize

import (
	"fmt"

	_ "github.com/mbobakov/grpc-consul-resolver" // It's important  负载均衡
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"mall-api/goods-web/global"
	"mall-api/goods-web/proto"
)

// InitSrvConn 初始化grpc服务  初始化user conn  新版本,并接入负载均衡
func InitSrvConn() {
	userConn, err := grpc.Dial(
		fmt.Sprintf("consul://%s:%d/%s?wait=14s",
			global.ServerConfig.ConsulInfo.Host, global.ServerConfig.ConsulInfo.Port, global.ServerConfig.GoodsSrvConfig.Name),
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`), //均匀负载(轮询)
	)
	if err != nil {
		zap.S().Fatal("Init Srv Conn With get user Conn error:", err.Error())
	}

	goodsSrvCli := proto.NewGoodsClient(userConn)
	global.GoodsSrvClient = goodsSrvCli
}
