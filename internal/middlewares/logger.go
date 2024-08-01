package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"go-sample/internal/constant"
)

// Logger 日志格式化中间件
func Logger() fiber.Handler {
	return logger.New(logger.Config{
		Format:        "${time} | ${method} | ${path} | ${status} | ${resBody}  | ${latency} | ${ip} | ${error} \n",
		TimeFormat:    constant.DefaultTimeFormat,
		Output:        nil,
		DisableColors: false,
	})
}