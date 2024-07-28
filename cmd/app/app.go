// @Title app.go
// @Description $END$
// @Author Zero - 2024/7/25 21:54:49

package app

import (
	"go-sample/configs"
	"go-sample/internal/engine"
	logs "go-sample/internal/logger"
	"go-sample/internal/server"
)

// Run App 服务启动入口
func Run() {
	// 初始化日志组件
	logs.Setup()
	// 初始化服务配置
	configs.Setup()
	// 初始化 HTTP 引擎
	handler := engine.SetupGinEngine()
	// 启动服务
	server.StartUp(handler)
}

