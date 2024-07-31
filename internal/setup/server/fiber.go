package server

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"go-sample/api"
	"go-sample/internal/constant"
	"go-sample/internal/global"
	"go-sample/internal/middlewares"
	logs "go-sample/internal/setup/logger"
	"os"
	"os/signal"
	"syscall"
)


// 创建并初始化 fiber 服务
func setup() *fiber.App {
	server := global.Conf.Server
	app := fiber.New(fiber.Config{
		AppName: server.Name,
		ServerHeader: "Go-Sample",
		// 请求正文大小限制
		BodyLimit: 1024 * 1024 * 8,
		// 路由区分大小写
		CaseSensitive: true,
		// 最大连接并发数
		Concurrency: 1024 * 256,
		// 打印路由信息
		EnablePrintRoutes: true,
		ColorScheme: fiber.DefaultColors,

		// 错误处理器
		ErrorHandler: middlewares.RecoveryHandler,

		// JSON 编解码，默认使用标准库，可替换成性能更高的库
		JSONDecoder: json.Unmarshal,
		JSONEncoder: json.Marshal,
		// 关闭输出调试信息
		DisableStartupMessage: true,
	})

	// 日志中间件
	app.Use(logger.New(logger.Config{
		Done:          nil,
		CustomTags:    nil,
		Format:        "${time} | ${method} | ${path} | ${status} | ${latency} | ${ip} | ${error}\n",
		TimeFormat:    constant.DefaultTimeFormat,
	}))

	// 注册 API 路由
	for _, m := range api.Modules {
		m.Init()
		m.Router(app)
	}
	return app
}

// Startup 启动HTTP 服务
func Startup() {
	// 创建服务
	server := setup()
	stop := make(chan os.Signal)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	stopped := make(chan struct{})
	// 异步启动 HTTP Server
	_ = global.Pool.Submit(func() {
		if err := server.Listen(global.Conf.Server.Addr()); err != nil {
			logs.Logger.FatalSf("Error on server startup: %s", err.Error())
		}
		close(stopped)
	})
	// 等待关闭信号
	<- stop
	logs.Logger.Info("start shutting down services.....")
	// 关闭 HTTP Server
	if err := server.Shutdown(); err != nil {
		logs.Logger.ErrorSf("failed to shutdown server: %s", err.Error())
	}
	// 等待关闭完成.
	<- stopped
	logs.Logger.Info("HTTP server closed success.")
}