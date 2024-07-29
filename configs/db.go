package configs

import "time"

// DataBase 数据库配置.
type DataBase struct {
	Host         string        `yaml:"host"`
	Port         string        `yaml:"port"`
	Username     string        `yaml:"username"`
	Password     string        `yaml:"password"`
	DbName       string        `yaml:"dbName"`
	Debug        bool          `yaml:"debug"`
	SlowSql      time.Duration `yaml:"slowSql"`
	CreateTables bool          `yaml:"createTables"`
	Pool         DPool         `yaml:"pool"`
}

// DPool 数据库连接池参数
type DPool struct {
	MaxOpenConn int           `yaml:"maxOpenConn"`
	MaxIdleConn int           `yaml:"maxIdleConn"`
	MaxLifeTime time.Duration `yaml:"maxLifeTime"`
}
