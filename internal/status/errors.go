package status

import "fmt"

// SysError 服务业务错误
type SysError struct {
	// 错误码
	Code int8 `json:"code"`
	// 错误消息
	Message string `json:"message"`
}

func (be *SysError) Error() string {
	return fmt.Sprintf("Bussiness Error [%d]: %s", be.Code, be.Message)
}

// 构建新的状态
func ofErr(code int8, msg string) *SysError {
	return &SysError{code, msg}
}