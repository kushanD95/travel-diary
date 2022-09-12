package main

import (
	"context"
	"fmt"
	"time"

	"github.com/kushanD95/traval-diary/app/controller"
	fiberconfig "github.com/kushanD95/traval-diary/package/config/fiber"
)

func main() {
	app := fiberconfig.SetupFiberApp()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	defer func(ctx context.Context) {
		ctx, cancel = context.WithTimeout(ctx, time.Second*10)
		defer cancel()
	}(ctx)

	controller.Controller(app)
	app.Listen(":" + "9000")
	fmt.Println("Application start")
}
