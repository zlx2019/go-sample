package server

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"go-sample/api"
	"go-sample/internal/constant"
	"go-sample/internal/global"
	"go-sample/internal/middlewares"
	logs "go-sample/internal/setup/logger"
	"os"
	"os/signal"
	"syscall"
	"time"
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

		// 服务错误统一处理器
		ErrorHandler: middlewares.GlobalErrorHandler,

		// JSON 编解码，默认使用标准库，可替换成性能更高的库
		JSONDecoder: json.Unmarshal,
		JSONEncoder: json.Marshal,
		// 关闭输出调试信息
		DisableStartupMessage: false,
	})

	// 注册中间件
	app.Use(middlewares.Cors(), recover.New())

	// 日志中间件
	app.Use(logger.New(logger.Config{
		Done:          nil,
		CustomTags:    nil,
		Format:        "${time} | ${method} | ${path} | ${status} | ${latency} | ${ip} | ${error}\n",
		TimeFormat:    constant.DefaultTimeFormat,
	}))


	// 注册 所有模块API
	for _, m := range api.Modules {
		m.Init()
		m.Router(app)
	}
	return app
}

// Startup 启动HTTP 服务
func Startup() {
	server := setup()
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	// 异步启动 HTTP Server
	_ = global.Pool.Submit(func() {
		if err := server.Listen(global.Conf.Server.Addr()); err != nil {
			logs.Logger.FatalSf("Error on server startup: %s", err.Error())
		}
	})
	// 等待关闭信号
	<- stop
	logs.Logger.Info("start shutting down services...")
	// 关闭 HTTP Server
	if err := server.ShutdownWithTimeout(time.Second *3); err != nil {
		logs.Logger.ErrorSf("failed to shutdown server: %s", err.Error())
	}
	logs.Logger.Info("HTTP server closed success.")
}