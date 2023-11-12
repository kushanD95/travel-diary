package validator

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/kushanD95/traval-diary/app/controller/dto"
	"github.com/kushanD95/traval-diary/package/config"
	commonDto "github.com/kushanD95/traval-diary/package/dto"
	"github.com/kushanD95/traval-diary/package/utils"
	"go.uber.org/zap"
)

func EnvConfigReq(ctx *fiber.Ctx) (*dto.EnvConfigReq, *commonDto.ErrorResponse) {
	lg := config.AppConfigutarion.GetLogger()
	lgFields := []zap.Field{zap.String(utils.METHOD, utils.ENV_CONFIG_REQ)}
	lg.Info(fmt.Sprintf(utils.ENV_CONFIG_REQ_LOG, utils.STARTED), lgFields...)
	var (
		request dto.EnvConfigReq
	)
	parseErr := ctx.BodyParser(&request)
	if parseErr != nil {
		errRes := &commonDto.ErrorResponse{
			Message: utils.INVALID_REQUEST_BODY,
			Code:    utils.StatusCode[utils.BadRequest],
			Error:   fmt.Sprintf("%v", parseErr.Error()),
		}

		lgFields = append(lgFields, zap.Any(utils.ERROR, parseErr))
		lg.Error(fmt.Sprintf(utils.ENV_CONFIG_REQ_LOG, utils.END_WITH_ERROR), lgFields...)
		return nil, errRes
	}
	validateErr := validate.Struct(&request)
	if validateErr != nil {
		errRes := &commonDto.ErrorResponse{
			Message: utils.INVALID_REQUEST_BODY,
			Code:    utils.StatusCode[utils.BadRequest],
			Error:   fmt.Sprintf("%v", validateErr.Error()),
		}
		lgFields = append(lgFields, zap.Any(utils.ERROR, validateErr))
		lg.Error(fmt.Sprintf(utils.ENV_CONFIG_REQ_LOG, utils.END_WITH_ERROR), lgFields...)
		return nil, errRes
	}

	lg.Info(fmt.Sprintf(utils.ENV_CONFIG_REQ_LOG, utils.END), lgFields...)
	return &request, nil
}
