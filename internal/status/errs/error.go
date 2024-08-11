package errs

import (
	"errors"
	"fmt"
)

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

// Wrap 将业务错误 和标准错误包装，合并为错误链.
func Wrap(e *Error, err error) error {
	return errors.Join(e, err)
}

func Is(src error, dst error) bool {
	return errors.Is(src, dst)
}