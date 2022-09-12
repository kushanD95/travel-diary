package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/kushanD95/traval-diary/app/response/builder"
)

func Check(ctx *fiber.Ctx) (err error) {
	fmt.Println("Check Handler")
	responseBuilder := builder.Response{
		Ctx:    ctx,
		Status: 200,
	}

	responseBuilder.BuildAndReturnResponse()
	return nil
}
