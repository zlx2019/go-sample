package configs

// Redis 缓存配置
type Redis struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
	Pool     RPool  `yaml:"pool"`
}

// RPool 连接池配置
type RPool struct {
}
