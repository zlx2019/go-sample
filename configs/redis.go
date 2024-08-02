package configs

import (
	"fmt"
	"time"
)

// Redis 缓存配置
type Redis struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
	Pool     RPool  `yaml:"pool"`
}

// RPool 连接池配置
type RPool struct {
	// 连接池大小
	MaxSize int `yaml:"maxSize"`
	// 最小空闲连接数,启动时就建立好的连接数量
	MinIdleConnNum int `yaml:"minIdleConnNum"`
	// 最大空闲连接数，多余的连接超时后关闭
	MaxIdleConnNum  int           `yaml:"maxIdleConnNum"`
	// 连接空闲超时时间
	ConnMaxIdleTime time.Duration `yaml:"connMaxIdleTime"`
}

func (r *Redis) Addr() string {
	return fmt.Sprintf("%s:%d", r.Host, r.Port)
}
