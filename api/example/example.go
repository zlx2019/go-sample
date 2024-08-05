package example

import "go-sample/internal/service"

// ApiExample Api
type ApiExample struct {
	serv *service.ExampleService
}

// NewExample 提供者
func NewExample(serv *service.ExampleService) *ApiExample {
	return &ApiExample{serv: serv}
}

// Init 初始化
func (e *ApiExample) Init() {
}
