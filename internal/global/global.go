package global

import (
	"github.com/panjf2000/ants/v2"
	"github.com/spf13/viper"
	"go-sample/configs"
	"gorm.io/gorm"
)

// 全局数据
var (
	// Conf 服务配置对象
	Conf *configs.Config
	Viper *viper.Viper

	// Dc 数据库客户端
	Dc *gorm.DB
	// Pool 协程池
	Pool *ants.Pool
)
