// @Title app.go
// @Description $END$
// @Author Zero - 2024/7/25 21:54:49

package app

import (
	"go-sample/configs"
	"go-sample/internal/logger"
)

// Run App 服务启动入口
func Run() {
	// 初始化日志组件
	logs.Init()
	// 加载配置文件.
	configs.Viper()
	logs.Logger.Info("Hello")
	logs.Logger.Debug("Hello")
	logs.Logger.Error("Hello")
}
