// @Title pool.go
// @Description 协程池
// @Author Zero - 2024/7/27 23:28:37

package pool

import (
	"github.com/panjf2000/ants/v2"
	"go-sample/configs"
	"go-sample/internal/global"
	logs "go-sample/internal/setup/logger"
	"time"
)

var pool *ants.Pool

// Setup 初始化协程池
func Setup() {
	p := global.Conf.Pool
	var err error
	// 创建池
	if pool, err = ants.NewPool(p.Size, withOptions(p)); err != nil {
		logs.Logger.Fatal("failed to init coroutine pool.")
	}
	global.Pool = pool
	logs.Logger.Info("[initialize pool success]")
}

// 配置.
func withOptions(p configs.Pool) func(*ants.Options) {
	return func(opt *ants.Options) {
		opt.DisablePurge = p.DisablePurge
		opt.ExpiryDuration = p.ExpiryDuration
		opt.PreAlloc = p.PreAlloc
		opt.Nonblocking = p.Nonblocking
		opt.MaxBlockingTasks = p.MaxBlockingTasks
		opt.PanicHandler = func(val interface{}) {
			// 池中任务发生 Panic
			logs.Logger.ErrorSf("pool task panic: %v", val)
		}
	}
}

// CleanUp 清理释放协程池
func CleanUp() {
	// 等待所有协程任务结束 or 超时
	if err := pool.ReleaseTimeout(time.Second * 10); err != nil {
		logs.Logger.ErrorSf("cleanup coroutine pool timeout: %s", err.Error())
	} else {
		logs.Logger.Info("[Cleanup coroutine pool complete]")
	}
}
