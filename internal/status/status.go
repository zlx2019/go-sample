package status

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

// Response 统一响应
type Response struct {
	Code int8    `json:"code"`
	Msg  string `json:"message"`
	Data any    `json:"data"`
}

// Ok 成功响应
func Ok(ctx *fiber.Ctx, data ...any) error {
	response := Response{
		Code: ok.Code,
		Msg:  ok.Message,
	}
	if len(data) > 0 {
		response.Data = data[0]
	}
	return writeTo(ctx, http.StatusOK, response)
}

// Fail 失败响应
func Fail(ctx *fiber.Ctx, err error) error {
	response := Response{}
	var se *SysError
	if errors.As(err, &se) {
		response.Code = se.Code
		response.Msg = se.Message
	}else {
		response.Code = fail.Code
		response.Msg = fail.Message
	}
	return writeTo(ctx, http.StatusInternalServerError, response)
}

// FailMsg 失败响应
func FailMsg(ctx *fiber.Ctx, messages ...string) error {
	response := Response{
		Code: fail.Code,
		Msg: fail.Message,
	}
	if len(messages) > 0 {
		response.Msg = messages[0]
	}
	return writeTo(ctx, http.StatusInternalServerError, response)
}


// 将响应流写回客户端.
func writeTo(ctx *fiber.Ctx, httpStatus int, response Response) error {
	//ctx.JSON(httpStatus, response)
	ctx.Status(http.StatusOK)
	return ctx.JSON(response)
}