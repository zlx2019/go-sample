package server

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"go-sample/internal/constant"
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

		// fiber config
		config := fiber.Config{
			AppName:      global.Conf.Server.Name,
			ServerHeader: "Go-Sample",
			// 请求正文大小限制
			BodyLimit: 1024 * 1024 * 8,
			// 路由区分大小写
			CaseSensitive: true,
			// 最大连接并发数
			Concurrency: 1024 * 256,
			// 打印路由信息
			EnablePrintRoutes: true,
			ColorScheme:       fiber.DefaultColors,

			// 服务错误统一处理器
			ErrorHandler: middlewares.GlobalErrorHandler,
			// JSON 编解码，默认使用标准库，可替换成性能更高的库
			JSONDecoder: json.Unmarshal,
			JSONEncoder: json.Marshal,
		}
		// 生产模式，不输出调试信息 && 路由信息
		if global.Conf.Server.Mode == constant.Release {
			config.DisableStartupMessage = true
			config.EnablePrintRoutes = false
		}
		server = fiber.New(config)
		// 注册中间件
		server.Use(recover.New(), middlewares.Cors(), middlewares.Logger())
		// 初始化所有模块 && 注册路由
		modules, _ := initModules(global.Dc, global.Rc)
		for _, m := range modules {
			m.Init()
			server.Route(m.Name(), m.Route(), m.Name())
		}

		server.Hooks().OnListen(func(listen fiber.ListenData) error {
			logs.Logger.InfoSf("\u001B[38;5;121mListening and serving HTTP on %s:%s\u001B[0m", listen.Host, listen.Port)
			return nil
		})
		server.Hooks().OnShutdown(func() error {
			logs.Logger.InfoSf("\u001B[1;91mHTTP Server Closed\u001B[0m")
			return nil
		})
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
	stop := make(chan os.Signal)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	// 等待信号
	sig := <-stop
	logs.Logger.InfoSf("server received signal: %v", sig)
	stopped := make(chan struct{})
	// 开始关闭服务
	go func() {
		if err := server.Shutdown(); err != nil {
			logs.Logger.ErrorSf("failed to shutdown server: %s", err.Error())
		}
		close(stopped)
	}()
	// 等待服务关闭完成.
	<-stopped
}
