package router

import (
	"github.com/gin-gonic/gin"
	"mall-api/user-web/api"
)

func InitBaseRouter(Router *gin.RouterGroup) {
	baseRouter := Router.Group("base")
	{
		baseRouter.GET("captcha", api.GetCaptcha)
		baseRouter.POST("send_sms", api.SendSms)
	}
}
