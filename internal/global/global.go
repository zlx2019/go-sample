package global

import (
	"github.com/spf13/viper"
	"go-sample/configs"
)

// 全局数据

var (
	// Conf 服务配置对象
	Conf *configs.Config
	Viper *viper.Viper
)
