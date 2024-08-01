// @Title router.go
// @Description ping api 路由
// @Author Zero - 2024/7/27 19:57:23

package ping

import (
	"github.com/gofiber/fiber/v2"
)


// Router Ping-模块路由
func (p *Ping) Router(app *fiber.App) {
	router := app.Group("/ping")
	router.Get("", p.Ping)
	router.Get("/pool", p.PoolStatus)
	router.Get("/db", p.DBStatus)
}