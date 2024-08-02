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

// RedisStatusReply Redis 连接统计信息
type RedisStatusReply struct {
	// 命中次数 (直接从池中取到连接的次数)
	Hits uint32 `json:"hits"`
	// 未命中次数 (从池中没有取到连接的次数)
	Misses uint32 `json:"misses"`
	// 尝试从池中获取连接，超时的次数
	Timeouts uint32	`json:"timeouts"`

	// 当前池中总连接数
	TotalConnNum uint32 `json:"total_conn_num" copier:"TotalConns"`
	// 当前池中空闲连接数
	IdleConnNum uint32 `json:"idle_conn_num" copier:"IdleConns"`
	// 共删除的空闲连接数
	RemoveConnNum uint32 `json:"stale_conn_num" copier:"StaleConns"`
}
