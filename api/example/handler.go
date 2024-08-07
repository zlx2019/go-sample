package example

import (
	"github.com/gofiber/fiber/v2"
	"go-sample/internal/status"
)

// Example 路由处理

func (e *Example) Hello(ctx *fiber.Ctx) error {
	return status.Ok(ctx, "hello world")
}