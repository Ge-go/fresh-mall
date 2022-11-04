package config

type UserSrvConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
	Name string `mapstructure:"name"`
}

type ServerConfig struct {
	Name        string        `mapstructure:"name"`
	Port        string        `mapstructure:"port"`
	UserSrvInfo UserSrvConfig `mapstructure:"user_srv"`
	JWTInfo     JWTConfig     `mapstructure:"jwt"`
	AliSmsInfo  AliSmsConfig  `mapstructure:"sms"`
	RedisInfo   RedisConfig   `mapstructure:"redis"`
	ConsulInfo  ConsulConfig  `mapstructure:"consul"`
}

type AliSmsConfig struct {
	ApiKey    string `mapstructure:"key"`
	ApiSecret string `mapstructure:"secret"`
	RegionId  string `mapstructure:"regionId"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	PassWord string `mapstructure:"password"`
}

type JWTConfig struct {
	SigningKey string `mapstructure:"sign_key"`
	PublicKey  string `mapstructure:"pub_key"`
}

type ConsulConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}
