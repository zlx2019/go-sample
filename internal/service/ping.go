package service

import (
	"github.com/redis/go-redis/v9"
	"go-sample/internal/store"
)

// PingService Ping 业务层
type PingService struct {
	store *store.PingStore
	rc    *redis.Client
}

// NewPingService 提供者
func NewPingService(store *store.PingStore, client *redis.Client) *PingService {
	return &PingService{store: store, rc: client}
}

// RedisStatus 查看 Redis 连接池信息
func (ps *PingService) RedisStatus() *redis.PoolStats {
	return ps.rc.PoolStats()
}
