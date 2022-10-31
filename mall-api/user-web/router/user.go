package router

import (
	"github.com/gin-gonic/gin"
	"mall-api/user-web/api"
)

func InitUserRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("user")
	{
		userRouter.GET("list", api.GetUserList)
	}
}
