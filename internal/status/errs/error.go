package errs

import "fmt"

// Error 服务业务错误
type Error struct {
	// 错误码
	Code int `json:"code"`
	// 错误消息
	Message string `json:"message"`
}

func (be *Error) Error() string {
	return fmt.Sprintf("Bussiness Error [%d]: %s", be.Code, be.Message)
}

// 构建错误
func of(code int, msg string) *Error {
	return &Error{code, msg}
}