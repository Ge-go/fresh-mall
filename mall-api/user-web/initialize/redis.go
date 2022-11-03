package initialize

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"

	"mall-api/user-web/global"
)

// InitRdsClient redis client init
func InitRdsClient() {
	global.RedisClient = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", global.ServerConfig.RedisInfo.Host, global.ServerConfig.RedisInfo.Port),
	})

	_, err := global.RedisClient.Ping(context.Background()).Result()
	if err != nil {
		panic(err.Error())
	}
}
