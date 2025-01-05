package api

import (
	"github.com/google/wire"
	"go-sample/api/email"
	"go-sample/api/example"
	"go-sample/api/ping"
	"go-sample/internal/service"
	"go-sample/internal/store"
)

// 所有模块 API 提供者
var (
	ExampleProvider = wire.NewSet(example.NewExample, service.NewExampleService, store.NewExampleStore)
	PingProvider    = wire.NewSet(ping.NewPing, service.NewPingService, store.NewPingStore)
	EmailProvider   = wire.NewSet(email.NewEmail, service.NewEmailService)
)
