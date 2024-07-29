package configs

import "fmt"

// Server 服务基础配置
type Server struct {
	Name      string `yaml:"name"`
	Host      string `yaml:"host"`
	Port      int    `yaml:"port"`
	Mode      string `yaml:"mode"`
	ApiPrefix string `yaml:"apiPrefix"`
}

func (s Server) Addr() string {
	if s.Host == "" || s.Port == 0 {
		return "0.0.0.0:8080"
	}
	return fmt.Sprintf("%s:%d", s.Host, s.Port)
}
