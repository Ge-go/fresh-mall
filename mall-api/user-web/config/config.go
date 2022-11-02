package config

type UserSrvConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type ServerConfig struct {
	Name        string        `mapstructure:"name"`
	Port        string        `mapstructure:"port"`
	UserSrvInfo UserSrvConfig `mapstructure:"user_srv"`
	JWTInfo     JWTConfig     `mapstructure:"jwt"`
}

type JWTConfig struct {
	SigningKey string `mapstructure:"sign_key"`
	PublicKey  string `mapstructure:"pub_key"`
}
