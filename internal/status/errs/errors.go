package errs

import "net/http"

// FailErr 系统内部错误，未知的
var (
	FailErr = of(fail, http.StatusText(http.StatusInternalServerError))
)

// 4xx 客户端错误
var (
	InvalidRequest   = of(40001, "无效的请求")
	Unauthenticated  = of(40002, "未认证")
	PermissionDenied = of(40003, "没有权限")
	NotFound         = of(40004, "资源不存在")
	AlreadyExists    = of(40009, "资源已存在，不可重复执行")
	RequestRateLimit = of(40029, "请求过于繁忙")
)

// 5xx 服务端内部错误
var (
	ServerInternal    = FailErr
	DataBaseError     = of(50002, "数据库发生错误")
	ServerUnavailable = of(50003, "服务器当前不可用")
)
