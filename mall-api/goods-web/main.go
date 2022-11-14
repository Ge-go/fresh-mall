package main

import (
	"fmt"
	"go.uber.org/zap"
	"mall-api/goods-web/global"
	"mall-api/goods-web/initialize"
	"mall-api/goods-web/utils/register/consul"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// 初始化router
	engine := initialize.Routers()

	// 初始化logger
	initialize.InitLogger()

	// 初始化yaml信息  常规项目依然使用viper + yaml即可
	//initialize.InitConfig()

	// 初始化nacos
	initialize.InitNacosConfig()

	// nacos to server config
	initialize.InitNacosToConfig()

	// 初始化validator,内置自定义mobile过滤
	if err := initialize.InitTrans("zh"); err != nil {
		panic(err)
	}

	// init user-srv conn
	initialize.InitSrvConn()

	// register to consul (inter)
	serviceID := fmt.Sprintf("%s:%d",
		global.ServerConfig.Name, global.ServerConfig.Port)
	rgsCli := consul.NewRegistryClient(global.ServerConfig.ConsulInfo.Host, global.ServerConfig.ConsulInfo.Port)
	err := rgsCli.Register(global.ServerConfig.Host, global.ServerConfig.Port,
		global.ServerConfig.Name, global.ServerConfig.Tags, serviceID)
	if err != nil {
		panic(err)
	}

	go func() {
		if err = engine.Run(fmt.Sprintf(":%d", global.ServerConfig.Port)); err != nil {
			panic(err)
		}
	}()

	//ctrl C 注销服务
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	if err = rgsCli.DeRegister(serviceID); err != nil {
		zap.S().Infof("注销失败 err:%v,serviceId:%s", err.Error(), serviceID)
	}
	zap.S().Info("deregister consul successful")
}
