package builder

import (
	"github.com/gofiber/fiber/v2"
)

type ErrorResponse struct {
	Message string
	Error   string
	Code    int
}

type Response struct {
	_        struct{}
	Ctx      *fiber.Ctx
	Payload  interface{}
	Status   int
	ErrorRes ErrorResponse
}

func (res *Response) BuildAndReturnResponse() {

	switch res.Status {
	case 200:
		res.Ctx.Status(res.Status).JSON(res.Payload)
	case 400:
	case 404:
		res.ErrorRes.Code = res.Status
		res.Ctx.Status(res.Status).JSON(res.ErrorRes)
	}
}
