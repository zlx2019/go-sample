// @Title gin.go
// @Author Zero - 2024/7/28 10:07:33

package engine

import (
	"github.com/gin-gonic/gin"
	"go-sample/api"
	"go-sample/configs"
	logs "go-sample/internal/logger"
	"net/http"
)

// SetupGinEngine 创建并初始化 Gin 服务引擎
func SetupGinEngine() http.Handler{
	engine := gin.New()
	gin.SetMode(configs.C.Server.Mode)
	// 注册中间件
	engine.Use(gin.Logger(), gin.Recovery())
	// 初始化API路由
	for _, m := range api.Modules {
		logs.Logger.InfoSf("Init module: [%s]", m.GetName())
		m.Init()
		m.Router(engine.Group("/" + m.GetName()))
	}
	return engine
}