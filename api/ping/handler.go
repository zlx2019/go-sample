package ping

import (
	"github.com/gofiber/fiber/v2"
	"go-sample/internal/global"
	"go-sample/internal/status"
	"go-sample/internal/status/errs"
	"go-sample/internal/tool/clone"
	"runtime"
)

// Ping 服务状态探测
func (p *ApiPing) ping(ctx *fiber.Ctx) error {
	return status.Ok(ctx, "OK")
}

// PoolStatus 查看服务协程池状态信息
func (p *ApiPing) poolStatus(ctx  *fiber.Ctx) error {
	pool := global.Pool
	reply := PoolStatusReply{
		IsClosed:      pool.IsClosed(),
		Capacity:      pool.Cap(),
		SysCoroutines: runtime.NumGoroutine(),
		Running:       pool.Running(),
		Waiting:       pool.Waiting(),
		Available:     pool.Free(),
	}
	return status.Ok(ctx, reply)
}

// 查看 Redis 连接信息
func (p *ApiPing) redisStatus(ctx *fiber.Ctx) error {
	state, err := clone.Clone[RedisStatusReply](p.serv.RedisStatus())
	if err != nil {
		return errs.FailErr
	}
	return status.Ok(ctx, state)
}

// DBStatus 查看数据库状态信息
func (p *ApiPing) dbStatus(ctx *fiber.Ctx) error {
	return nil
}