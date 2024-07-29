package logs

import (
	"github.com/zlx2019/spoor"
	"go-sample/internal/constant"
	"go.uber.org/zap"
)

var Logger *spoor.Spoor

// Setup 初始化日志组件
func Setup() {
	logger, err := spoor.NewSpoor(&spoor.Config{
		LogDir:             "./logs",
		FileName:           "go-sample",
		Level:           	zap.DebugLevel,
		WriteFile:      	true,
		FileSeparate: 		false,
		Plugins:            []zap.Option{zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel)},
		WrapSkip:           1,
		LogTimeFormat: 		constant.DefaultTimeFormat,
	}, spoor.WithFileSizeCutter(10, 30, constant.MB * 50))
	if err != nil {
		panic(err)
	}
	Logger = logger
}
