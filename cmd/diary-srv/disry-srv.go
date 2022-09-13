package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/kushanD95/traval-diary/app/controller"
	"github.com/kushanD95/traval-diary/package/config"
	fiberconfig "github.com/kushanD95/traval-diary/package/config/fiber"
)

var appConfig *config.AppConfig

func init() {
	appConfig = &config.AppConfig{}
}

func main() {
	app := fiberconfig.SetupFiberApp()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	defer func(ctx context.Context) {
		ctx, cancel = context.WithTimeout(ctx, time.Second*10)
		defer cancel()
	}(ctx)

	db := appConfig.SetupDB()
	if db == nil {
		log.Fatal("db connection failed")
	}

	controller.Controller(app)
	app.Listen(":" + "9000")
	fmt.Println("Application start")
}
