package global

import (
	ut "github.com/go-playground/universal-translator"
	"mall-api/goods-web/proto"

	"mall-api/goods-web/config"
)

var (
	ServerConfig   *config.ServerConfig
	Trans          ut.Translator
	NacosConfig    *config.NacosConfig
	GoodsSrvClient proto.GoodsClient
)
