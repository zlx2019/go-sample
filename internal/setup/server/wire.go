//go:build wireinject
// +build wireinject

package server

import (
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
	"go-sample/api"
	"go-sample/api/example"
	"go-sample/api/ping"
	"gorm.io/gorm"
)

// 模块依赖注入，获取API路由列表
func initModules(*gorm.DB,*redis.Client) ([]api.Module, error) {
	wire.Build(
		api.ExampleProvider,
		api.PingProvider,
		ProvideModules)
	return nil,nil
}

func ProvideModules(exam *example.Example, ping *ping.Ping) []api.Module {
	return []api.Module{
		exam,
		ping,
	}
}