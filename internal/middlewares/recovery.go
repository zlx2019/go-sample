package middlewares

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"go-sample/internal/status"
	"net"
	"net/http"
	"os"
	"strings"
)

// GlobalErrorHandler 统一错误响应处理
func GlobalErrorHandler(c *fiber.Ctx, err error) error {
	response := status.OfErr(err)
	return c.Status(fiber.StatusInternalServerError).JSON(response)
}

// Recovery 自定义 Panic 捕获处理(也可以直接使用官方提供的)
func Recovery() fiber.Handler {
	return func(ctx *fiber.Ctx) (err error) {
		defer func() {
			if info := recover(); info != nil {
				// 检查连接
				var brokenPipe bool
				if ne, ok := info.(*net.OpError); ok {
					var se *os.SyscallError
					if errors.As(ne, &se) {
						seStr := strings.ToLower(se.Error())
						if strings.Contains(seStr, "broken pipe") ||
							strings.Contains(seStr, "connection reset by peer") {
							brokenPipe = true
						}
					}
				}
				if brokenPipe {
					// 连接不可写会
					err = errors.New(http.StatusText(fiber.StatusInternalServerError))
					return
				}
				if e,ok := info.(error); ok {
					_ = status.FailWithErr(ctx, e)
					err = e
				}
				err = errors.New(http.StatusText(fiber.StatusInternalServerError))
				return
			}
		}()
		return ctx.Next()
	}
}