package fiberconfig

import "github.com/gofiber/fiber/v2"

func SetupFiberApp() *fiber.App {
	fiberApp := fiber.New()

	return fiberApp
}

func ShutdownFiberApp(app *fiber.App) (err error) {
	err = app.Shutdown()

	return err
}
