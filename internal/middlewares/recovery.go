package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"go-sample/internal/status"
)

// RecoveryHandler Fiber引擎 Panic 错误处理
func RecoveryHandler(c *fiber.Ctx, err error) error {
	response := status.OfErr(err)
	return c.Status(fiber.StatusInternalServerError).JSON(response)
}


// Recovery 错误处理中间件
//func Recovery() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		defer func() {
//			// 捕获 panic
//			if err := recover(); err != nil {
//
//				// 检查连接是否易损
//				var brokenPipe bool
//				if ne, ok := err.(*net.OpError); ok {
//					var se *os.SyscallError
//					if errors.As(ne, &se) {
//						seStr := strings.ToLower(se.Error())
//						if strings.Contains(seStr, "broken pipe") ||
//							strings.Contains(seStr, "connection reset by peer") {
//							brokenPipe = true
//						}
//					}
//				}
//				// 获取 HTTP Request
//				request, _ := httputil.DumpRequest(c.Request, false)
//				if brokenPipe {
//					// 连接已无法使用，记录日志并直接结束.
//					zap.L().Error(c.Request.URL.Path, zap.Any("error", err), zap.String("request", string(request)))
//					c.Error(err.(error)) //nolint: errcheck
//					c.Abort()
//					return
//				}
//				// 记录日志
//				if e, ok := err.(error); ok {
//					logs.Logger.Sugar().
//						WithOptions(zap.AddCallerSkip(1)).
//						Errorf("handler request [%s] failed, caused by: %v", c.Request.URL.Path, e)
//				}
//				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
//					"code":    1,
//					"message": "系统错误.",
//				})
//			}
//		}()
//		c.Next()
//	}
//}