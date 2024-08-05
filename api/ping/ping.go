// @Title ping.go
// @Description ping api
// @Author Zero - 2024/7/27 19:57:17

package ping

import "go-sample/internal/service"

// ApiPing API
type ApiPing struct {
	serv *service.PingService
}

// NewPing ApiPing 提供者
func NewPing(serv *service.PingService) *ApiPing {
	return &ApiPing{serv: serv}
}

// Init ApiPing-模块初始化
func (p *ApiPing) Init() {
}
