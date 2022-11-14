package initialize

import (
	"github.com/gin-gonic/gin"
	"mall-api/goods-web/middleware"

	"mall-api/goods-web/router"
)

func Routers() *gin.Engine {
	engine := gin.Default()
	// 配置跨域插件
	engine.Use(middleware.Cors())

	apiGroup := engine.Group("/g/v1")
	{
		// goods server
		router.InitGoodsRouter(apiGroup)
	}

	return engine
}
