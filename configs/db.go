package configs

// DataBase 数据库配置.
type DataBase struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DbName   string `yaml:"db_name"`
	Pool     DPool  `yaml:"pool"`
}

// DPool 数据库连接池参数
type DPool struct {
}
