package srv_config

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"go.uber.org/zap"
	"mall_srvs/user_srv/global"
)

// RegisterToConsul srv register to consul
func RegisterToConsul() {
	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d",
		global.ServerConfig.ConsulInfo.Host, global.ServerConfig.ConsulInfo.Port)

	client, err := api.NewClient(cfg)
	if err != nil {
		zap.S().Errorw("new client consul err", "msg", err.Error())
		panic(err)
	}

	serviceID := fmt.Sprintf("%s:%d",
		global.ServerConfig.Name, global.ServerConfig.Port)
	// 生成注册对象
	err = client.Agent().ServiceRegister(&api.AgentServiceRegistration{
		Name:    global.ServerConfig.Name,
		ID:      serviceID,
		Port:    global.ServerConfig.Port,
		Address: global.ServerConfig.Host,
		Tags:    []string{"wS", "sW", "test"},
		Check: &api.AgentServiceCheck{
			GRPC: fmt.Sprintf("%s:%d",
				global.ServerConfig.Host, global.ServerConfig.Port),
			Timeout:                        "5s",
			Interval:                       "5s",
			DeregisterCriticalServiceAfter: "15s",
		},
	})

	if err != nil {
		zap.S().Errorw("register consul err", "msg", err.Error())
		panic(err)
	}
	// 优雅退出
	//go func() {
	//	quit := make(chan os.Signal)
	//	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	//	<-quit
	//	if err = client.Agent().ServiceDeregister(serviceID); err != nil {
	//		zap.S().Info("deregister consul failed")
	//	}
	//	zap.S().Info("deregister consul successful")
	//}()
}
