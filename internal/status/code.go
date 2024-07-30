package status

// 错误码

// OK 请求成功 or 请求错误
var (
	ok = ofErr(0, "ok")
)

// 请求失败
var (
	fail = ofErr(1, "failed")
)
