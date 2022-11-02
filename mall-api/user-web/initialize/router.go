package initialize

import (
	"github.com/gin-gonic/gin"
	"mall-api/user-web/middleware"

	"mall-api/user-web/router"
)

func Routers() *gin.Engine {
	engine := gin.Default()
	// 配置跨域插件
	engine.Use(middleware.Cors())

	apiGroup := engine.Group("/u/v1")
	{
		// user server
		router.InitUserRouter(apiGroup)
		// base server
		router.InitBaseRouter(apiGroup)
	}

	return engine
}
