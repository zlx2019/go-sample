package configs

// Server 服务基础配置
type Server struct {
	Name      string `yaml:"name"`
	Host      string `yaml:"host"`
	Port      int    `yaml:"port"`
	Mode      string `yaml:"mode"`
	ApiPrefix string `yaml:"api_prefix"`
}
