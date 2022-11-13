package config

type ServerConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
	Name string `mapstructure:"name" json:"name"`

	MySQLInfo  MySQLConfig  `mapstructure:"mysql" json:"mysql"`
	ConsulInfo ConsulConfig `mapstructure:"consul" json:"consul"`
	EsInfo     EsConfig     `mapstructure:"es" json:"es"`
}

type EsConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
}

type MySQLConfig struct {
	Name     string `mapstructure:"name" json:"name"`
	PassWord string `mapstructure:"pwd" json:"pwd"`
	IP       string `mapstructure:"ip" json:"ip"'`
	Port     int    `mapstructure:"port" json:"port"`
	DB       string `mapstructure:"db" json:"db"`
}

type ConsulConfig struct {
	Host string   `mapstructure:"host" json:"host"`
	Port int      `mapstructure:"port" json:"port"`
	Tags []string `json:"tags"`
}
