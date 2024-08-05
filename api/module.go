package api

// @Title       module.go
// @Description API 模块抽象
// @Author      Zero.
// @Create      2024-08-05 13:58

import (
	"github.com/gofiber/fiber/v2"
)

// Module API基层模块
type Module interface {
	Init()
	Route(app *fiber.App)
}