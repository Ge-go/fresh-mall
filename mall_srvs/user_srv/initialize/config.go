package initialize

import (
	"encoding/json"

	"github.com/fsnotify/fsnotify"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"mall_srvs/user_srv/global"
	"mall_srvs/user_srv/utils"
)

func GetEnvInfo(env string) bool {
	viper.AutomaticEnv()
	return viper.GetBool(env)
}

func InitNacosToConfig() {
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr: global.NacosConfig.Nacos.Host,
			Port:   uint64(global.NacosConfig.Nacos.Port),
		},
	}

	cc := constant.ClientConfig{
		NamespaceId:         global.NacosConfig.Nacos.Namespace,
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "tmp/nacos/log",
		CacheDir:            "tmp/nacos/cache",
		LogLevel:            "debug",
	}

	client, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfig":  cc,
	})
	if err != nil {
		panic(err)
	}

	cfg, err := client.GetConfig(vo.ConfigParam{
		DataId: global.NacosConfig.Nacos.DataId,
		Group:  global.NacosConfig.Nacos.Group,
	})
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal([]byte(cfg), &global.ServerConfig)
	if err != nil {
		zap.S().Errorw("json unmarshal server config err", "msg", err.Error())
		panic(err)
	}

	//todo 后期是否要移除  随机端口
	port, _ := utils.GetFreePort()
	global.ServerConfig.Port = port
}

// InitConfig read Config
func InitConfig() {
	configFileName := "user_srv/config-debug.yaml" // debug环境
	pro := GetEnvInfo("FRESH_MALL_PRO")            // 生产环境
	if pro {
		configFileName = "user_srv/config-pro.yaml" //生产环境
	} else if GetEnvInfo("FRESH_TEST") { //test环境 for windows
		configFileName = "user_srv/config-test.yaml"
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

	//todo 后期是否要移除
	port, _ := utils.GetFreePort()
	global.ServerConfig.Port = port

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
