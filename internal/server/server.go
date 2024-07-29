// @Title server.go
// @Author Zero - 2024/7/27 20:06:07

package server

import (
	"context"
	"errors"
	"go-sample/configs"
	logs "go-sample/internal/logger"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)


// StartUp 服务
func StartUp(handler http.Handler) {
	// HTTP 服务配置
	server := &http.Server{
		Addr:        configs.C.Server.Addr(),
		Handler:     handler,
		ReadTimeout: time.Second * 30,
		WriteTimeout: time.Second * 30,
		MaxHeaderBytes: 1 << 20,
	}
	defer server.Close()

	// 优雅关闭服务
	stop := make(chan os.Signal)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	stopped := make(chan struct{})

	// 异步启动HTTP Server
	go func() {
		logs.Logger.InfoSf("Listening and serving HTTP on %s", server.Addr)
		if err := server.ListenAndServe(); err != nil{
			if !errors.Is(err, http.ErrServerClosed) {
				logs.Logger.FatalSf("Error on server startup: ",err.Error())
			}
			// 服务被主动关闭
			stopped <- struct{}{}
		}
	}()
	// 阻塞进程，直到收到信号
	<- stop
	logs.Logger.Info("Start shutting down services...")

	// 停止 HTTP 服务
	if err := server.Shutdown(context.Background()); err != nil {
		logs.Logger.ErrorSf("Failed to shut down service.")
	}
	// 等待 HTTP 服务终止
	<- stopped
	logs.Logger.Info("HTTP server closed")
}
