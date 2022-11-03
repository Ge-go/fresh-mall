package global

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-redis/redis/v8"

	"mall-api/user-web/config"
)

var (
	ServerConfig *config.ServerConfig
	Trans        ut.Translator
	RedisClient  *redis.Client
)