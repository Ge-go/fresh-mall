package config

type GoodsSrvConfig struct {
	Host string `mapstructure:"host" json:"host"'`
	Port int    `mapstructure:"port" json:"port"`
	Name string `mapstructure:"name" json:"name"`
}

type ServerConfig struct {
	Name           string         `mapstructure:"name" json:"name"`
	Host           string         `mapstructure:"host" json:"host"`
	Port           int            `mapstructure:"port" json:"port"`
	Tags           []string       `mapstructure:"tags" json:"tags"`
	JWTInfo        JWTConfig      `mapstructure:"jwt" json:"jwt"`
	ConsulInfo     ConsulConfig   `mapstructure:"consul" json:"consul"`
	GoodsSrvConfig GoodsSrvConfig `mapstructure:"goods_srv" json:"goods_srv"`
}

type JWTConfig struct {
	SigningKey string `mapstructure:"sign_key" json:"sign_key"`
	PublicKey  string `mapstructure:"pub_key" json:"pub_key"`
}

type ConsulConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
}
