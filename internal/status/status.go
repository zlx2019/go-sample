package status

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Response 统一响应
type Response struct {
	Code int8    `json:"code"`
	Msg  string `json:"message"`
	Data any    `json:"data"`
}

// Ok 成功响应
func Ok(ctx *gin.Context, data ...any) {
	response := Response{
		Code: ok.Code,
		Msg:  ok.Message,
	}
	if len(data) > 0 {
		response.Data = data[0]
	}
	writeTo(ctx, http.StatusOK, response)
}

// Fail 失败响应
func Fail(ctx *gin.Context, err error) {
	response := Response{}
	var se *SysError
	if errors.As(err, &se) {
		response.Code = se.Code
		response.Msg = se.Message
	}else {
		response.Code = fail.Code
		response.Msg = fail.Message
	}
	writeTo(ctx, http.StatusInternalServerError, response)
}

// FailMsg 失败响应
func FailMsg(ctx *gin.Context, messages ...string) () {
	response := Response{
		Code: fail.Code,
		Msg: fail.Message,
	}
	if len(messages) > 0 {
		response.Msg = messages[0]
	}
	writeTo(ctx, http.StatusInternalServerError, response)
}


// 将响应流写回客户端.
func writeTo(ctx *gin.Context, httpStatus int, response Response) {
	ctx.JSON(httpStatus, response)
}