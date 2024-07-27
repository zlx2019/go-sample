// @Title ping.go
// @Description ping api
// @Author Zero - 2024/7/27 19:57:17

package ping

// Ping API
type Ping struct {}

func (p *Ping) GetName() string {
	return "ping"
}

// Init Ping API 初始化
func (p *Ping) Init() {
}
