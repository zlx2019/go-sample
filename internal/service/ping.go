package service

import (
	"github.com/redis/go-redis/v9"
	"go-sample/internal/dao"
)

// PingService Ping 业务层
type PingService struct {
	repo *dao.PingRepo
	rc *redis.Client
}

// NewPingService 提供者
func NewPingService(repo *dao.PingRepo, client *redis.Client) *PingService {
	return &PingService{repo: repo, rc: client}
}

// RedisStatus 查看 Redis 连接池信息
func (ps *PingService) RedisStatus() *redis.PoolStats {
	return ps.rc.PoolStats()
}