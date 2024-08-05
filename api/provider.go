package api

import (
	"github.com/google/wire"
	"go-sample/api/example"
	"go-sample/api/ping"
	"go-sample/internal/dao"
	"go-sample/internal/service"
)

// 所有模块 API 提供者
var (
	ExampleProvider = wire.NewSet(example.NewExample, service.NewExampleService, dao.NewExampleRepo)
	PingProvider = wire.NewSet(ping.NewPing, service.NewPingService, dao.NewPingRepo)
)