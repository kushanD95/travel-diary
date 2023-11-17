package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kushanD95/traval-diary/app/controller/handler"
)

func Controller(app *fiber.App) {
	baseRoute := app.Group("/travel-diary/v1")

	baseRoute.Get("/check", handler.Check)
	baseRoute.Post("/ready", handler.Ready)
	baseRoute.Post("/login", handler.Login)
	baseRoute.Post("/register", handler.Register)
}
