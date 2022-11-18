package router

import (
	"github.com/gin-gonic/gin"
	"mall-api/goods-web/api/banners"
	"mall-api/goods-web/middleware"
)

func InitBannerRouter(Router *gin.RouterGroup) {
	BannerRouter := Router.Group("banners") //.Use(middleware.Trace())
	{
		BannerRouter.GET("", banners.List)                                                          // 轮播图列表页
		BannerRouter.DELETE("/:id", middleware.JWTAuth(), middleware.IsAdminAuth(), banners.Delete) // 删除轮播图
		BannerRouter.POST("", middleware.JWTAuth(), middleware.IsAdminAuth(), banners.New)          //新建轮播图
		BannerRouter.PUT("/:id", middleware.JWTAuth(), middleware.IsAdminAuth(), banners.Update)    //修改轮播图信息
	}
}
