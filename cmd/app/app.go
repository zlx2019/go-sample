// @Title app.go
// @Description $END$
// @Author Zero - 2024/7/25 21:54:49

package app

import (
	"go-sample/configs"
	"go-sample/internal/global"
	"go-sample/internal/setup/database"
	"go-sample/internal/setup/logger"
	"go-sample/internal/setup/pool"
	"go-sample/internal/setup/server"
)

// Startup App 服务启动入口
func Startup() {
	// ===================== Startup ===============================
	// 初始化日志组件
	logs.Setup()
	// 初始化配置
	global.Conf, global.Viper = configs.Setup()
	// 初始化数据库.
	database.Setup()
	// 初始化协程池
	pool.Setup()
	// 初始化并启动 HTTP 服务
	server.Startup()

	// Running...

	// ===================== CleanUp ===============================
	// 释放数据库连接池
	database.CleanUp()
	// 释放协程池
	pool.CleanUp()
	// 刷新日志缓冲
	logs.Cleanup()
}
