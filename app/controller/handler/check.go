package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/kushanD95/traval-diary/app/response/builder"
	"github.com/kushanD95/traval-diary/package/config"
	"github.com/kushanD95/traval-diary/package/utils"
)

func Check(ctx *fiber.Ctx) (err error) {
	lg := config.AppConfigutarion.GetLogger()
	lg.Info(fmt.Sprintf(utils.CHECK_HANDLER, utils.STARTED))
	responseBuilder := builder.Response{
		Ctx:     ctx,
		Payload: Success,
		Status:  utils.StatusCode[Success],
	}

	responseBuilder.BuildAndReturnResponse()

	lg.Info(fmt.Sprintf(utils.CHECK_HANDLER, utils.END))
	return nil
}
