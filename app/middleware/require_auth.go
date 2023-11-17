package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func RequireAuth(ctx *fiber.Ctx) {
	fmt.Println("____________middleare______________")
	ctx.Next()
}
