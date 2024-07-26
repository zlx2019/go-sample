// @Title app.go
// @Description $END$
// @Author Zero - 2024/7/25 21:54:49

package app

import (
	"go-sample/configs"
	logs "go-sample/internal/logger"
)

// Run App 服务启动入口
func Run() {
	// 初始化日志组件
	logs.Setup()
	// 加载配置文件.
	configs.Setup()
	logs.Logger.Info("Hello")
}
