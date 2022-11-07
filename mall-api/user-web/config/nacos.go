package config

// NacosConfig Nacos nacos config
type NacosConfig struct {
	Nacos Nacos `yaml:"nacos"`
}

type Nacos struct {
	Host      string `yaml:"host"`
	Port      int    `yaml:"port"`
	Namespace string `yaml:"namespace"`
	User      string `yaml:"user"`
	Password  string `yaml:"password"`
	DataId    string `yaml:"dataid"`
	Group     string `yaml:"group"`
}
