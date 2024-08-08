package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"go-sample/internal/constant"
	"io"
	"os"
)

// Logger 日志格式化中间件
func Logger() fiber.Handler {
	file, _ := os.OpenFile("./logs/fiber.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	log.SetLevel(log.LevelDebug)
	return logger.New(logger.Config{
		Format:        "${time} | ${method} | ${path} | ${status} | ${resBody}  | ${latency} | ${ip} | ${error} \n",
		TimeFormat:    constant.DefaultTimeFormat,
		Output:        io.MultiWriter(os.Stdout, file),
		DisableColors: false,
	})
}