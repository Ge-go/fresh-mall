package main

import (
	"mall-api/user-web/global"
	"mall-api/user-web/initialize"
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

	// 初始化redis
	initialize.InitRdsClient()

	// init user-srv conn
	initialize.InitSrvConn()

	if err := engine.Run(global.ServerConfig.Port); err != nil {
		panic(err)
	}
}
