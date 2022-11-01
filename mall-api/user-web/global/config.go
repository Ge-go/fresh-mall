package global

import (
	ut "github.com/go-playground/universal-translator"

	"mall-api/user-web/config"
)

var (
	ServerConfig *config.ServerConfig
	Trans        ut.Translator
)
