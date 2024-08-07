package middlewares

import (
	. "github.com/gofiber/fiber/v2"
	"strings"
)

// Cors Fiber 跨域中间件
func Cors() Handler {
	return func(c *Ctx) error {
		origin := string(c.Request().Header.Peek(HeaderOrigin))
		if origin != "" {
			c.Request().Header.Set(HeaderAccessControlAllowOrigin, "*")
			c.Request().Header.Set(HeaderAccessControlAllowMethods, strings.Join([]string{
				MethodGet,
				MethodPost,
				MethodHead,
				MethodOptions,
				MethodPut,
				MethodDelete,
				MethodPatch,
			}, ","))
			c.Request().Header.Set(HeaderAccessControlAllowHeaders, "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			c.Request().Header.Set(HeaderAccessControlExposeHeaders, "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Request().Header.Set(HeaderAccessControlAllowCredentials, "true")
		}
		if c.Method() == MethodOptions {
			return c.SendStatus(StatusNoContent)
		}
		return c.Next()
	}
}
