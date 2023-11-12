package handler

import (
	"fmt"

	"github.com/kushanD95/traval-diary/app/controller/dto"
	"github.com/kushanD95/traval-diary/app/controller/handler/validator"
	"go.uber.org/zap"

	"github.com/gofiber/fiber/v2"
	"github.com/kushanD95/traval-diary/app/response/builder"
	"github.com/kushanD95/traval-diary/app/services"
	"github.com/kushanD95/traval-diary/package/config"
	commonDto "github.com/kushanD95/traval-diary/package/dto"
	"github.com/kushanD95/traval-diary/package/utils"
)

func Check(ctx *fiber.Ctx) (err error) {
	lg := config.AppConfigutarion.GetLogger()
	lg.Info(fmt.Sprintf(utils.CHECK_HANDLER, utils.STARTED))
	responseBuilder := builder.Response{
		Ctx:     ctx,
		Payload: utils.Success,
		Status:  utils.StatusCode[utils.Success],
	}

	responseBuilder.BuildAndReturnResponse()

	lg.Info(fmt.Sprintf(utils.CHECK_HANDLER, utils.END))
	return nil
}

func Ready(ctx *fiber.Ctx) (err error) {
	lg := config.AppConfigutarion.GetLogger()
	lgFields := []zap.Field{zap.String(utils.METHOD, utils.READY)}
	lg.Info(fmt.Sprintf(utils.READY_HANDLER, utils.STARTED), lgFields...)

	var (
		request         *dto.EnvConfigReq
		response        *commonDto.EnvConfig
		errRes          *commonDto.ErrorResponse
		responseBuilder *builder.Response
		statusCode      int
	)

	defer func() {

		if err := recover(); err != nil {
			responseBuilder := builder.Response{
				Ctx: ctx,
				ErrorRes: &commonDto.ErrorResponse{
					Message: utils.INTERNAL_SERVER_ERROR,
					Code:    utils.StatusCode[utils.InternalServer],
					Error:   fmt.Sprintf("%v", err),
				},
				Status: utils.StatusCode[utils.InternalServer],
			}

			responseBuilder.BuildAndReturnResponse()
			lgFields = append(lgFields, zap.Any(utils.ERROR, err))
			lg.Error(fmt.Sprintf(utils.READY_HANDLER, utils.END_WITH_ERROR), lgFields...)
		}
	}()

	request, errRes = validator.EnvConfigReq(ctx)
	if errRes == nil {
		service := services.CreateConfigService(nil)
		response, errRes = service.FetchConfigService(request)
		statusCode = utils.StatusCode[utils.Success]
	}

	if errRes != nil {
		statusCode = errRes.Code
	}

	responseBuilder = &builder.Response{
		Ctx:      ctx,
		ErrorRes: errRes,
		Payload:  response,
		Status:   statusCode,
	}

	responseBuilder.BuildAndReturnResponse()

	lg.Info(fmt.Sprintf(utils.READY_HANDLER, utils.END), lgFields...)
	return nil
}
