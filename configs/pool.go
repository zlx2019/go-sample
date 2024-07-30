package configs

import "time"

// Pool 协程池配置.
//
// Size 池的最大容量
// DisablePurge 是否关闭回收空闲的work
// ExpiryDuration  DisablePurge为false时才生效, 如5 * time.Second 表示空闲5秒后的work会被回收掉
// PreAlloc 初始化时预先分配
// Nonblocking 指定是否使用非阻塞模式执行任务。如果设置为true，则在协程池已满的情况下，任务会立即返回一个err，而不是等待空闲协程。
// MaxBlockingTasks 阻塞模式下,最多允许阻塞等待的协程数量
type Pool struct {
	Size             int           `yaml:"size"`
	DisablePurge     bool          `yaml:"disablePurge"`
	ExpiryDuration   time.Duration `yaml:"expiryDuration"`
	PreAlloc         bool          `yaml:"preAlloc"`
	Nonblocking      bool          `yaml:"nonblocking"`
	MaxBlockingTasks int           `yaml:"maxBlockingTasks"`
}
