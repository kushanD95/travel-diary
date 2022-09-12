package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/kushanD95/traval-diary/app/response/builder"
	"github.com/kushanD95/traval-diary/package/utils"
)

func Check(ctx *fiber.Ctx) (err error) {
	fmt.Println("Check Handler")
	responseBuilder := builder.Response{
		Ctx:     ctx,
		Payload: Success,
		Status:  utils.StatusCode[Success],
	}

	responseBuilder.BuildAndReturnResponse()
	return nil
}
