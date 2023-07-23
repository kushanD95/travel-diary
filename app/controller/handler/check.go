package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kushanD95/traval-diary/app/response/builder"
	"github.com/kushanD95/traval-diary/package/config"
	"github.com/kushanD95/traval-diary/package/utils"
)

func Check(ctx *fiber.Ctx) (err error) {
	lg := config.AppConfigutarion.GetLogger()
	lg.Info("Check Handler Started")
	responseBuilder := builder.Response{
		Ctx:     ctx,
		Payload: Success,
		Status:  utils.StatusCode[Success],
	}

	responseBuilder.BuildAndReturnResponse()

	lg.Info("Check Handler End")
	return nil
}
