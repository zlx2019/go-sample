// @Title ping.go
// @Description ping api
// @Author Zero - 2024/7/27 19:57:17

package ping

import "go-sample/internal/service"

const name = "ping"

// Ping 服务状态探测模块
type Ping struct {
	serv *service.PingService
}

// Name Ping 模块名 && 路由前缀
func (p *Ping) Name() string {
	return name
}

// NewPing Ping 提供者
func NewPing(serv *service.PingService) *Ping {
	return &Ping{serv: serv}
}

// Init Ping-模块初始化
func (p *Ping) Init() {
}
