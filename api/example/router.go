package example

import (
	"github.com/gofiber/fiber/v2"
	"go-sample/internal/status"
)

// Route of Example
func (e *Example) Route(app *fiber.App) {
	router := app.Group("example")
	router.Get("", func(ctx *fiber.Ctx) error {
		return status.Ok(ctx, "hello world")
	})
}


