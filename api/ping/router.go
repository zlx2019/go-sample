// @Title router.go
// @Description ping api 路由
// @Author Zero - 2024/7/27 19:57:23

package ping

import (
	"github.com/gin-gonic/gin"
)

// Router Ping
func (p *Ping) Router(router *gin.RouterGroup) {
	router.GET("", p.Ping)
	router.GET("/pool", p.PoolStatus)
	router.GET("/db", p.DBStatus)
}