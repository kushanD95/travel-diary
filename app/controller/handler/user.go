package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/kushanD95/traval-diary/app/response/builder"
	"github.com/kushanD95/traval-diary/app/services"
	"github.com/kushanD95/traval-diary/package/config"
	"github.com/kushanD95/traval-diary/package/dto"
	"github.com/kushanD95/traval-diary/package/utils"
	"go.uber.org/zap"
)

func Register(ctx *fiber.Ctx) error {
	lg := config.AppConfigutarion.GetLogger()
	lgFields := []zap.Field{zap.String(utils.METHOD, utils.REGISTER)}
	lg.Info(fmt.Sprintf(utils.REGISTER_HANDLER, utils.STARTED), lgFields...)

	var user dto.User
	defer func() {
		if err := recover(); err != nil {
			responseBuilder := builder.Response{
				Ctx: ctx,
				ErrorRes: &dto.ErrorResponse{
					Message: utils.INTERNAL_SERVER_ERROR,
					Code:    utils.StatusCode[utils.InternalServer],
					Error:   fmt.Sprintf("%v", err),
				},
				Status: utils.StatusCode[utils.InternalServer],
			}

			responseBuilder.BuildAndReturnResponse()
			lgFields = append(lgFields, zap.Any(utils.ERROR, err))
			lg.Info(fmt.Sprintf(utils.REGISTER_HANDLER, utils.END_WITH_ERROR), lgFields...)
		}
	}()

	if err := ctx.BodyParser(&user); err != nil {
		lgFields = append(lgFields, zap.Any(utils.ERROR, err))
		lg.Error(utils.BODY_PARSER_ERROR, lgFields...)
		responseBuilder := builder.Response{
			Ctx: ctx,
			ErrorRes: &dto.ErrorResponse{
				Message: utils.BadRequest,
				Code:    utils.StatusCode[utils.BadRequest],
				Error:   fmt.Sprintf("%v", err),
			},
			Status: utils.StatusCode[utils.BadRequest],
		}

		responseBuilder.BuildAndReturnResponse()
		lg.Info(fmt.Sprintf(utils.REGISTER_HANDLER, utils.END_WITH_ERROR), lgFields...)
		return nil
	}
	services.Register(&user)
	lg.Info(fmt.Sprintf(utils.RECEIVED_PAYLOAD, user), lgFields...)
	lg.Info(fmt.Sprintf(utils.REGISTER_HANDLER, utils.END), lgFields...)
	return nil
}

func Login(ctx *fiber.Ctx) error {
	lg := config.AppConfigutarion.GetLogger()
	lgFields := []zap.Field{zap.String(utils.METHOD, utils.LOGIN)}
	lg.Info(fmt.Sprintf(utils.LOGIN_HANDLER, utils.STARTED), lgFields...)

	lg.Info(fmt.Sprintf(utils.LOGIN_HANDLER, utils.END), lgFields...)
	return nil
}
