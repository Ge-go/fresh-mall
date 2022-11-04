package initialize

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"mall_srvs/user_srv/global"
	"mall_srvs/user_srv/utils"
)

func GetEnvInfo(env string) bool {
	viper.AutomaticEnv()
	return viper.GetBool(env)
}

// InitConfig read Config
func InitConfig() {
	configFileName := "user_srv/config-debug.yaml" // debug环境
	pro := GetEnvInfo("FRESH_MALL_PRO")            // 生产环境
	if pro {
		configFileName = "user_srv/config-pro.yaml" //生产环境
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

	if pro { //如果是生产环境,交给consul
		port, _ := utils.GetFreePort()
		global.ServerConfig.Port = port
	}

	// 监控yaml
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		v.ReadInConfig()
		v.Unmarshal(&global.ServerConfig)
		zap.S().Infof("the user server config is changed:%v", global.ServerConfig)
	})
}
