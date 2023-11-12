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
	lgFields := []zap.Field{zap.String("Method", "Register")}
	lg.Info(fmt.Sprintf(utils.REGISTER_HANDLER, utils.STARTED))

	var user dto.User
	defer func() {
		if err := recover(); err != nil {
			responseBuilder := builder.Response{
				Ctx: ctx,
				ErrorRes: &dto.ErrorResponse{
					Message: "Internal server error",
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
		lg.Info(fmt.Sprintf(utils.REGISTER_HANDLER, utils.END_WITH_ERROR))
		return nil
	}
	services.Register(&user)
	lg.Info(fmt.Sprintf("received payload %v", user))
	lg.Info(fmt.Sprintf(utils.REGISTER_HANDLER, utils.END))
	return nil
}

func Login(ctx *fiber.Ctx) error {
	lg := config.AppConfigutarion.GetLogger()
	lg.Info(fmt.Sprintf(utils.LOGIN_HANDLER, utils.STARTED))

	lg.Info(fmt.Sprintf(utils.LOGIN_HANDLER, utils.END))
	return nil
}
