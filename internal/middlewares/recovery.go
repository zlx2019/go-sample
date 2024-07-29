package middlewares

import (
	"errors"
	"github.com/gin-gonic/gin"
	logs "go-sample/internal/setup/logger"
	"go.uber.org/zap"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
)

// Recovery 错误处理中间件
func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			// 捕获 panic
			if err := recover(); err != nil {

				// 检查连接是否易损
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					var se *os.SyscallError
					if errors.As(ne, &se) {
						seStr := strings.ToLower(se.Error())
						if strings.Contains(seStr, "broken pipe") ||
							strings.Contains(seStr, "connection reset by peer") {
							brokenPipe = true
						}
					}
				}
				// 获取 HTTP Request
				request, _ := httputil.DumpRequest(c.Request, false)

				// 连接已无法写回数据
				if brokenPipe {
					zap.L().Error(c.Request.URL.Path, zap.Any("error", err), zap.String("request", string(request)))
					c.Error(err.(error)) //nolint: errcheck
					c.Abort()
					return
				}
				// 记录日志
				logs.Logger.ErrorSf("Request Path: [%s] err: %v", c.Request.URL.Path, err)
				// 返回500
				c.AbortWithStatusJSON(200, gin.H{
					"code":    1,
					"message": "系统错误.",
				})
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}