package builder

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kushanD95/traval-diary/package/dto"
)

type Response struct {
	_        struct{}
	Ctx      *fiber.Ctx
	Payload  interface{}
	Status   int
	ErrorRes *dto.ErrorResponse
}

func (res *Response) BuildAndReturnResponse() {

	switch res.Status {
	case 200:
		res.Ctx.Status(res.Status).JSON(res.Payload)
	default:
		res.Ctx.Status(res.Status).JSON(&res.ErrorRes)
	}
}
