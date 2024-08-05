package example

import (
	"github.com/gofiber/fiber/v2"
)

// Route ApiExample 路由映射
func (e *ApiExample) Route(app *fiber.App) {
	router := app.Group("example")
	router.Get("", e.Hello)
}


