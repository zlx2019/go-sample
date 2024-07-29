package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Cors Gin 请求跨域处理中间件
func Cors() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		method := ctx.Request.Method
		// 放行所有 OPTIONS 请求
		if method == "OPTIONS" {
			ctx.AbortWithStatus(http.StatusNoContent)
		}
		origin := ctx.Request.Header.Get("Origin") //请求头部
		if origin != "" {
			ctx.Header("Access-Control-Allow-Origin", "*")  // 可将将 * 替换为指定的域名
			ctx.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			ctx.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			ctx.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			ctx.Header("Access-Control-Allow-Credentials", "true")
		}
		// 放行
		ctx.Next()
	}
}
