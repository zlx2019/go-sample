package cache

import (
	"context"
	"github.com/redis/go-redis/v9"
	"go-sample/internal/global"
	logs "go-sample/internal/setup/logger"
	"sync"
)

var once sync.Once
var _client *redis.Client

// Setup 初始化并连接 Redis
func Setup() {
	once.Do(func() {
		c := global.Conf.Redis
		pool := c.Pool
		_client = redis.NewClient(&redis.Options{
			Addr:                       c.Addr(),
			ClientName:                 global.Conf.Server.Name,
			OnConnect: onConnectHook,
			Username:                   c.Username,
			Password:                   c.Password,
			DB:                         c.DB,
			ContextTimeoutEnabled:      true,
			PoolSize:                   pool.MaxSize,
			MinIdleConns:               pool.MinIdleConnNum,
			MaxIdleConns:               pool.MaxIdleConnNum,
			ConnMaxIdleTime:            pool.ConnMaxIdleTime,
		})
		// ping
		pong, err := _client.Ping(context.Background()).Result()
		if err != nil {
			logs.Logger.FatalSf("redis connect fail err: %v", err)
		}
		logs.Logger.InfoSf("[connected database success: %s]", pong)
		global.Rc = _client
	})
}

// 连接 hook
func onConnectHook(ctx context.Context, conn *redis.Conn) error {
	if err := conn.Ping(ctx).Err(); err != nil {
		logs.Logger.ErrorSf("New redis conn fail to ping: %v", err)
		return err
	}
	logs.Logger.DebugSf("New redis conn... current connects: %d", _client.PoolStats().TotalConns)
	return nil
}

// CleanUp 释放Redis连接
func CleanUp() {
	_ = _client.Close()
	logs.Logger.Info("【 Cleanup redis complete 】")
}