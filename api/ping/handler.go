package ping

import (
	"github.com/gofiber/fiber/v2"
	"go-sample/internal/global"
	"go-sample/internal/status"
	"runtime"
)

// Ping 服务状态探测
func (p *Ping) Ping(ctx *fiber.Ctx) error {
	return status.Ok(ctx)
}

// PoolStatus 查看服务协程池状态信息
func (p *Ping) PoolStatus(ctx  *fiber.Ctx) error {
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

// DBStatus 查看数据库状态信息
func (p *Ping) DBStatus(ctx *fiber.Ctx) error {
	return nil
}