package configs

var C = Config{}

// Config 项目总配置
type Config struct {
	Server Server   `yaml:"server"`
	DB     DataBase `yaml:"db"`
	Redis  Redis    `yaml:"cache"`
	Pool   Pool     `yaml:"pool"`
}

