// @Title app.go
// @Description $END$
// @Author Zero - 2024/7/25 21:54:49

package app

import (
	"go-sample/configs"
	"go-sample/internal/global"
	"go-sample/internal/setup/database"
	"go-sample/internal/setup/logger"
	"go-sample/internal/setup/server"
)

// Run App 服务启动入口
func Run() {
	// 初始化日志组件
	logs.Setup()
	// 初始化服务配置
	global.Conf, global.Viper = configs.Setup()
	// 初始化数据库.
	database.Setup()
	// 初始化并启动 HTTP 服务
	server.StartUp()
}
