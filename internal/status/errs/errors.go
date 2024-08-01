package errs

import "net/http"

// FailErr 系统内部错误，未知的
var (
	FailErr = of(fail, http.StatusText(http.StatusInternalServerError))
)

// 4xx 客户端错误
var (
	InvalidRequest = of(40001, "Invalid Request") // 无效的请求
	NotFound       = of(40004, "Not found")       // 目标不存在
)

// 5xx 服务端内部错误
var (
	ServerInternal = FailErr
	DataBaseError       = of(50002, "DataBase Error") // 数据库错误
)
