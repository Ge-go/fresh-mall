package router

import (
	"github.com/gin-gonic/gin"
	"mall-api/goods-web/middleware"

	"mall-api/goods-web/api/goods"
)

func InitGoodsRouter(Router *gin.RouterGroup) {
	goodsRouter := Router.Group("goods").Use(middleware.Trace())
	{
		goodsRouter.GET("list", goods.List)                                                      //获取用户列表
		goodsRouter.POST("new", middleware.JWTAuth(), middleware.IsAdminAuth(), goods.New)       //新建商品
		goodsRouter.GET("/:id", goods.Detail)                                                    // 商品详情
		goodsRouter.DELETE("/:id", middleware.JWTAuth(), middleware.IsAdminAuth(), goods.Delete) //删除商品

		goodsRouter.GET("/:id/stocks", goods.Stocks)                                                  // 获取商品库存
		goodsRouter.PATCH("/:id", middleware.JWTAuth(), middleware.IsAdminAuth(), goods.UpdateStatus) // 更新商品状态
		goodsRouter.PUT("/:id", middleware.JWTAuth(), middleware.IsAdminAuth(), goods.Update)         // 更新商品
	}
}
