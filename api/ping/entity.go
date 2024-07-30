package ping

type RequestPing struct {
}

// PoolStatusReply 协程池状态响应
type PoolStatusReply struct {
	// 是否已关闭
	IsClosed bool `json:"is_closed"`
	// 允许开启的最大协程数
	Capacity int `json:"capacity"`
	// 当前进程开启的协程总数
	SysCoroutines int `json:"sys_coroutines"`
	// 池中正在运行的协程数量
	Running int `json:"running"`
	// 任务队列中的任务数量
	Waiting int `json:"waiting"`
	// 当前可用协程数量
	Available int `json:"available"`
}
