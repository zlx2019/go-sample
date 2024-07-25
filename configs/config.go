package configs


var C = Config{}

// Config 服务配置
type Config struct {
	Server Server   `yaml:"server"`
	DB     DataBase `yaml:"db"`
	Redis  Redis    `yaml:"redis"`
}
