package ping

import (
	"github.com/gin-gonic/gin"
	"go-sample/internal/global"
	"go-sample/internal/status"
	"runtime"
)

// Ping 服务状态探测
func (p *Ping) Ping(ctx *gin.Context) {
	status.Ok(ctx)
}

// PoolStatus 查看服务协程池状态信息
func (p *Ping) PoolStatus(ctx *gin.Context)  {
	pool := global.Pool
	reply := PoolStatusReply{
		IsClosed:      pool.IsClosed(),
		Capacity:      pool.Cap(),
		SysCoroutines: runtime.NumGoroutine(),
		Running:       pool.Running(),
		Waiting:       pool.Waiting(),
		Available:     pool.Free(),
	}
	status.Ok(ctx, reply)
}

// DBStatus 查看数据库状态信息
func (p *Ping) DBStatus(ctx *gin.Context) {
	
}