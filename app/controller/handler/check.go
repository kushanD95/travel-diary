package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Check(ctx *fiber.Ctx) (err error) {
	fmt.Println("Check Handler")
	ctx.Status(200)
	return nil
}
