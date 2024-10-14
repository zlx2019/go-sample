package email

import (
	"github.com/gofiber/fiber/v2"
)

func (e *Email) Route() func(fiber.Router) {
	return func(router fiber.Router) {
		// 获取 AccessToken
		router.Get("/accessToken", GetAccessToken)
	}
}
