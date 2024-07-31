package status

import (
	"errors"
	"github.com/gofiber/fiber/v2"
)


type R map[string]any

// Response 统一响应
type Response struct {
	Code int    `json:"code"`
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
	return writeTo(ctx, fiber.StatusOK, response)
}

// FailWithErr 失败响应
func FailWithErr(ctx *fiber.Ctx, err error) error {
	response := Response{}
	var se *SysError
	if errors.As(err, &se) {
		response.Code = se.Code
		response.Msg = se.Message
	}else {
		response.Code = fail.Code
		if err.Error() == "" {
			response.Msg = fail.Message
		}
	}
	return writeTo(ctx, fiber.StatusInternalServerError, response)
}

// Fail 失败响应
func Fail(ctx *fiber.Ctx, messages ...string) error {
	response := Response{
		Code: fail.Code,
		Msg: fail.Message,
	}
	if len(messages) > 0 {
		response.Msg = messages[0]
	}
	return writeTo(ctx, fiber.StatusInternalServerError, response)
}


// 将响应流写回客户端.
func writeTo(ctx *fiber.Ctx, status int, response Response) error {
	return ctx.Status(status).JSON(response)
}


// OfErr 根据Error构建响应信息
func OfErr(err error) Response {
	var se *SysError
	if errors.As(err, &se) {
		return Response{Code: se.Code, Msg: se.Message}
	}
	return Response{Code: fail.Code, Msg:  err.Error()}
}