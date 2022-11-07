package initialize

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"mall-api/user-web/global"
)

func InitNacosConfig() {
	v := viper.New()
	// read yaml
	configFileName := "user-web/config-nacos.yaml"
	v.SetConfigFile(configFileName)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := v.Unmarshal(&global.NacosConfig); err != nil {
		panic(err)
	}

	// 监控nacos yaml
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		v.ReadInConfig()
		v.Unmarshal(&global.NacosConfig)
		zap.S().Infof("the user server config is changed:%v", global.ServerConfig)
	})
}
