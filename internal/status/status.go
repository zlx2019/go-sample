package status

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"go-sample/internal/status/errs"
)

const (
	successCode = 0
	successMsg = "OK"
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
		Code: successCode,
		Msg:  successMsg,
	}
	if len(data) > 0 {
		response.Data = data[0]
	}
	return writeTo(ctx, fiber.StatusOK, response)
}

// Fail 失败响应
func Fail(ctx *fiber.Ctx, messages ...string) error {
	response := of(errs.FailErr)
	if len(messages) > 0 {
		response.Msg = messages[0]
	}
	return writeTo(ctx, fiber.StatusInternalServerError, response)
}

// FailWithErr 失败响应
func FailWithErr(ctx *fiber.Ctx, err error) error {
	response := OfErr(err)
	return writeTo(ctx, fiber.StatusInternalServerError, response)
}

// OfErr 根据未知错误，构建响应信息
func OfErr(err error) Response {
	var se *errs.Error
	var response Response
	response.Code = errs.FailErr.Code
	if errors.As(err, &se) {
		response.Code = se.Code
		response.Msg = se.Message
	} else {
		if err.Error() == "" {
			response.Msg = errs.FailErr.Message
		}else {
			response.Msg = err.Error()
		}
	}
	return response
}

func writeTo(ctx *fiber.Ctx, status int, response Response) error {
	return ctx.Status(status).JSON(response)
}

func of(err *errs.Error) Response {
	return Response{
		Code: err.Code,
		Msg:  err.Message,
	}
}

