package logs

import (
	"github.com/zlx2019/spoor"
	"go.uber.org/zap"
	"time"
)

var Logger *spoor.Spoor

// Setup 日志组件初始化
func Setup() {
	logger, err := spoor.NewSpoor(&spoor.Config{
		LogDir:             "./logs",
		FileName:           "go-sample",
		LogLevel:           zap.DebugLevel,
		LogWriterFile:      true,
		LogWriterFromLevel: false,
		MaxSaveTime:        time.Hour * 24 * 7,
		MaxFileSize:        100,
		Style:              false,
		Plugins:            []zap.Option{zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel)},
		WrapSkip:           1,
	})
	if err != nil {
		panic(err)
	}
	Logger = logger
}
