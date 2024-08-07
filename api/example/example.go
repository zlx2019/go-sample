package example

import "go-sample/internal/service"

const name = "example"

// Example Api
type Example struct {
	serv *service.ExampleService
}

// Name Example 模块名字 && 前缀
func (e *Example) Name() string {
	return name
}

// NewExample 提供者
func NewExample(serv *service.ExampleService) *Example {
	return &Example{serv: serv}
}

// Init 初始化
func (e *Example) Init() {
}
