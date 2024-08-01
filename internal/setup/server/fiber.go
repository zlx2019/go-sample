package server

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"go-sample/api"
	"go-sample/internal/global"
	"go-sample/internal/middlewares"
	logs "go-sample/internal/setup/logger"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var server *fiber.App
var once sync.Once

// 初始化 HTTP 服务
func setup() {
	once.Do(func() {
		server = fiber.New(fiber.Config{
			AppName: global.Conf.Server.Name,
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
		server.Use(recover.New(),middlewares.Cors(), middlewares.Logger())
		// 注册 所有模块API
		for _, m := range api.Modules {
			m.Init()
			m.Router(server)
		}
	})
}

// Startup 启动HTTP 服务
func Startup() {
	setup()
	go cleanup()
	if err := server.Listen(global.Conf.Server.Addr()); err != nil {
		logs.Logger.FatalSf("Error on server startup: %s", err.Error())
	}
}

// 通过信号关闭HTTP服务
func cleanup() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	// 等待信号
	sig := <- stop
	logs.Logger.InfoSf("start shutting down services, on siginal: %v", sig)
	stopped := make(chan struct{})
	// 开始关闭服务
	go func() {
		if err := server.Shutdown(); err != nil {
			logs.Logger.ErrorSf("failed to shutdown server: %s", err.Error())
		}
		close(stopped)
	}()
	// 等待服务关闭完成.
	<- stopped
	logs.Logger.Info("HTTP server closed success.")
}