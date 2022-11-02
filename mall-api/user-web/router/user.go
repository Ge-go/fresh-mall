package router

import (
	"github.com/gin-gonic/gin"

	"mall-api/user-web/api"
	"mall-api/user-web/middleware"
)

func InitUserRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("user")
	{
		userRouter.GET("list", middleware.JWTAuth(), middleware.IsAdminAuth(), api.GetUserList) //获取用户列表
		userRouter.POST("pwd_login", api.PassWordLogin)                                         // 根据账号密码进行登录
	}
}
