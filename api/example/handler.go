package example

import (
	"github.com/gofiber/fiber/v2"
	"go-sample/internal/status"
)

// ApiExample 路由处理

func (e *ApiExample) Hello(ctx *fiber.Ctx) error {
	return status.Ok(ctx, "hello world")
}