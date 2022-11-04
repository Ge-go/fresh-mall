package initialize

import (
	"fmt"
	"mall-api/user-web/proto"

	"github.com/hashicorp/consul/api"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"mall-api/user-web/global"
)

// InitSrvConn 初始化grpc服务  初始化user conn
func InitSrvConn() {
	// todo 这么从注册中心去服务发现,也有问题,就是我并没有接入新的注册中心的内容
	// 从注册中心获取用户服务的信息
	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d",
		global.ServerConfig.ConsulInfo.Host, global.ServerConfig.ConsulInfo.Port)
	consulClient, err := api.NewClient(cfg)
	if err != nil {
		zap.S().Errorw("[consul create client err]", "msg", err.Error())
		return
	}
	srvInfo, err := consulClient.Agent().ServicesWithFilter(fmt.Sprintf(`Service == "%s"`,
		global.ServerConfig.UserSrvInfo.Name))
	if err != nil {
		zap.S().Errorw("[consul agent get srv]", "msg", err.Error())
		return
	}

	userSrvHost := ""
	userSrvPort := 0
	for _, val := range srvInfo {
		userSrvHost = val.Address
		userSrvPort = val.Port
	}
	if userSrvHost == "" {
		zap.S().Errorw("[from consul get srv info err]")
		return
	}

	zap.S().Infof("i'm get user-srv service from consul,ip:%s;port:%d;",
		userSrvHost, userSrvPort)

	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", userSrvHost, userSrvPort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		zap.S().Errorw("[grpc.Dial] conn err", "msg", err.Error())
		return
	}

	// todo 1.用户服务下线了 2.改端口了 3.改ip了  整体逻辑需要优化
	// 好处,省的多次tcp请求
	// 问题,一个链接,多个协程共用,性能  这里最好使用连接池
	client := proto.NewUserClient(conn)

	global.UserSrvClient = client
}
