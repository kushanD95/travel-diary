package builder

import (
	"github.com/gofiber/fiber/v2"
)

type ErrorResponse struct {
	Code    int
	Message string
	Error   string
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
	default:
		res.Ctx.Status(res.Status).JSON(res.ErrorRes)
	}
}
