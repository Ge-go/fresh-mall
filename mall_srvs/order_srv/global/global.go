package global

import (
	"gorm.io/gorm"
	"mall_srvs/goods_srv/proto"
	pt "mall_srvs/inventory_srv/proto"

	"mall_srvs/order_srv/config"
)

var (
	DB           *gorm.DB
	ServerConfig *config.ServerConfig
	NacosConfig  *config.NacosConfig

	GoodsSrvClient     proto.GoodsClient
	InventorySrvClient pt.InventoryClient
)
