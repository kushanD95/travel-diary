package main

import (
	"context"
	"time"

	"github.com/kushanD95/traval-diary/app/controller"
	"github.com/kushanD95/traval-diary/app/services"
	"github.com/kushanD95/traval-diary/package/config"
	fiberconfig "github.com/kushanD95/traval-diary/package/config/fiber"
	"github.com/kushanD95/traval-diary/package/utils"
)

func init() {
	config.AppConfigutarion = &config.AppConfig{}
	config.AppConfigutarion.InitConfig()
	config.AppConfigutarion.InitLogger()
}

func main() {
	// setup fiber
	app := fiberconfig.SetupFiberApp()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	defer func(ctx context.Context) {
		ctx, cancel = context.WithTimeout(ctx, time.Second*10)
		defer cancel()
	}(ctx)

	//setup db connection
	db := config.AppConfigutarion.SetupDB()
	dbConn := services.DBConn{
		Db: db,
	}
	dbConn.SetupDB()

	//register the controller
	controller.Controller(app)
	app.Listen(utils.Colon + config.AppConfigutarion.AppPort)
}
