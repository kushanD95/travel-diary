package main

import (
	"context"
	"time"

	"github.com/kushanD95/traval-diary/app/controller"
	"github.com/kushanD95/traval-diary/package/config"
	fiberconfig "github.com/kushanD95/traval-diary/package/config/fiber"
	"github.com/kushanD95/traval-diary/package/utils"
)

var appConfig *config.AppConfig

func init() {
	appConfig = &config.AppConfig{}
	appConfig.InitConfig()
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
	_ = db

	controller.Controller(app)
	app.Listen(utils.Colon + appConfig.AppPort)
}
