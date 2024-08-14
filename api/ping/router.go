// @Title router.go
// @Description ping api 路由
// @Author Zero - 2024/7/27 19:57:23

package ping

import (
	"github.com/gofiber/fiber/v2"
	"go-sample/internal/status"
	"go-sample/internal/status/errs"
)

// Route on Ping module
func (p *Ping) Route() func(router fiber.Router) {
	return func(router fiber.Router) {
		router.Get("", p.ping)
		router.Get("/pool", p.poolStatus)
		router.Get("/db", p.dbStatus)
		router.Get("/redis", p.redisStatus)
		router.Post("/test", func(ctx *fiber.Ctx) error {
			var body map[string]any
			if err := ctx.BodyParser(&body); err != nil {
				return errs.FailErr
			}
			return status.Ok(ctx, body)
		})
	}
}
