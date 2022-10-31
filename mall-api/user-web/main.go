package main

import "mall-api/user-web/initialize"

func main() {
	// 初始化router
	engine := initialize.Routers()

	// 初始化logger
	initialize.InitLogger()

	if err := engine.Run(":8080"); err != nil {
		panic(err)
	}
}
