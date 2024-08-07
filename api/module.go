package api

import "github.com/gofiber/fiber/v2"

// @Title       module.go
// @Description API 模块抽象
// @Author      Zero.
// @Create      2024-08-05 13:58

// Module API基层模块
type Module interface {
	// Init 初始化函数
	Init()
	// Route 模块路由映射
	Route() func(route fiber.Router)
	// Name 模块名称
	Name() string
}
