package config

type ServerConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
	Name     string `mapstructure:"name"`

	MySQLInfo  MySQLConfig  `mapstructure:"mysql"`
	ConsulInfo ConsulConfig `mapstructure:"consul"`
}

type MySQLConfig struct {
	Name     string `mapstructure:"name"`
	PassWord string `mapstructure:"pwd"`
	IP       string `mapstructure:"ip"`
	Port     int    `mapstructure:"port"`
	DB       string `mapstructure:"db"`
}

type ConsulConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}
