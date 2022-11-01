package initialize

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"mall-api/user-web/global"
)

func GetEnvInfo(env string) bool {
	viper.AutomaticEnv()
	return viper.GetBool(env)
}

func InitConfig() {
	configFileName := "user-web/config-debug.yaml" // debug环境
	pro := GetEnvInfo("FRESH_MALL_PRO")            // 生产环境
	if pro {
		configFileName = "user-web/config-pro.yaml" //生产环境
	}

	v := viper.New()
	// read yaml
	v.SetConfigFile(configFileName)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := v.Unmarshal(&global.ServerConfig); err != nil {
		panic(err)
	}

	zap.S().Infof("the user server config is:%v", global.ServerConfig)

	// 监控yaml
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		v.ReadInConfig()
		v.Unmarshal(&global.ServerConfig)
		zap.S().Infof("the user server config is changed:%v", global.ServerConfig)
	})
}
