package router

import (
	"github.com/gin-gonic/gin"
	"mall-api/goods-web/middleware"

	"mall-api/goods-web/api/goods"
)

func InitGoodsRouter(Router *gin.RouterGroup) {
	goodsRouter := Router.Group("goods")
	{
		goodsRouter.GET("list", goods.List) //获取用户列表
		goodsRouter.POST("new", middleware.JWTAuth(), middleware.IsAdminAuth(), goods.New)
	}
}
