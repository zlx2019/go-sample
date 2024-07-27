// @Title router.go
// @Description ping api 路由
// @Author Zero - 2024/7/27 19:57:23

package ping

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Router Ping
func (p *Ping) Router(router *gin.RouterGroup) {
	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 0,
			"message": "OK",
		})
	})
}


