package example

import "github.com/gofiber/fiber/v2"


// Route on Example module
func (e *Example) Route() func(router fiber.Router){
	return func(router fiber.Router){
		router.Get("", e.Hello)
	}
}

