// @Title module.go
// @Description api 抽象模块
// @Author Zero - 2024/7/27 19:48:48

package api

import (
	"github.com/gin-gonic/gin"
	"go-sample/api/ping"
)

// Module API基层模块
type Module interface {
	GetName() string
	Init()
	Router(router *gin.RouterGroup)
}

// Modules 要加载的模块列表
var Modules []Module

// 注册模块
func registerModule(modules ...Module) {
	Modules = append(Modules, modules...)
}

// 初始化模块路由
func init() {
	registerModule(&ping.Ping{})
}