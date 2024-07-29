// @Title hertz.go
// @Author Zero - 2024/7/28 10:13:26

package engine
//import (
//	"github.com/cloudwego/hertz/pkg/app/middlewares/server/recovery"
//	"github.com/cloudwego/hertz/pkg/app/server"
//	"go-sample/api"
//	logs "go-sample/internal/logger"
//	"time"
//)
//
//// SetupHertzEngine 创建并初始化 Hertz 服务引擎
//func SetupHertzEngine() *server.Hertz {
//	engine := server.New(
//		// 开启热部署
//		server.WithAutoReloadRender(true, 0),
//		// 基础 path
//		server.WithBasePath("/hertz/"),
//		// 优雅停机等待时长
//		server.WithExitWaitTime(time.Second*5))
//	// 注册中间件
//	engine.Use(recovery.Recovery())
//	// 注册路由
//	for _, m := range api.Modules {
//		logs.Logger.InfoSf("Init module: [%s]", m.GetName())
//		m.Init()
//		m.Router2(engine.Group(m.GetName()))
//	}
//	return engine
//}
//