// @Title gin.go
// @Author Zero - 2024/7/27 20:06:07

package server

// 创建并初始化 Gin 服务引擎func setup() http.Handler {
////	engine := gin.New()
////	gin.SetMode(configs.C.Server.Mode)
////	gin.ForceConsoleColor()
////	// 注册中间件
////	engine.Use(middlewares.Cors(), middlewares.Recovery(),gin.Logger())
////	// 初始化API路由
////	for _, m := range api.Modules {
////		logs.Logger.InfoSf("Init module: [%s]", m.GetName())
////		m.Init()
////		m.Router(engine.Group("/" + m.GetName()))
////	}
////	return engine
////}
//

// Startup 启动 HTTP 服务
//func Startup() {
//	// HTTP 服务配置
//	server := &http.Server{
//		Addr:           global.Conf.Server.Addr(),
//		Handler:        setup(),
//		ReadTimeout:    time.Second * 30,
//		WriteTimeout:   time.Second * 30,
//		MaxHeaderBytes: 1 << 20,
//	}
//	defer server.Close()
//
//	// 优雅关闭服务
//	stop := make(chan os.Signal)
//	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
//	stopped := make(chan struct{})
//
//	// 异步启动HTTP Server
//	global.Pool.Submit(func() {
//		logs.Logger.InfoSf("Listening and serving HTTP on %s", server.Addr)
//		if err := server.ListenAndServe(); err != nil {
//			if !errors.Is(err, http.ErrServerClosed) {
//				logs.Logger.FatalSf("Error on server startup: %s", err.Error())
//			}
//			// 服务被主动关闭
//			close(stopped)
//		}
//	})
//	// 阻塞进程，直到收到信号
//	<-stop
//	logs.Logger.Info("Start shutting down services...")
//
//	// 停止 HTTP 服务
//	if err := server.Shutdown(context.Background()); err != nil {
//		logs.Logger.ErrorSf("Failed to shut down service.")
//	}
//	// 等待 HTTP 服务终止
//	<-stopped
//	logs.Logger.Info("HTTP server closed")
//}
