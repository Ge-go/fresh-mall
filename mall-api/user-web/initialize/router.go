package initialize

import (
	"github.com/gin-gonic/gin"
	"mall-api/user-web/router"
)

func Routers() *gin.Engine {
	engine := gin.Default()

	userGroup := engine.Group("/u/v1")
	{ // user server
		router.InitUserRouter(userGroup)
	}

	return engine
}
