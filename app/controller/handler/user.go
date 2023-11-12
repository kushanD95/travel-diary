package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/kushanD95/traval-diary/app/controller/dto"
	"github.com/kushanD95/traval-diary/app/controller/handler/validator"
	"github.com/kushanD95/traval-diary/app/response/builder"
	"github.com/kushanD95/traval-diary/app/services"
	"github.com/kushanD95/traval-diary/package/config"
	commondto "github.com/kushanD95/traval-diary/package/dto"
	"github.com/kushanD95/traval-diary/package/utils"
	"go.uber.org/zap"
)

func Register(ctx *fiber.Ctx) error {
	lg := config.AppConfigutarion.GetLogger()
	lgFields := []zap.Field{zap.String(utils.METHOD, utils.REGISTER)}
	lg.Info(fmt.Sprintf(utils.REGISTER_HANDLER, utils.STARTED), lgFields...)

	var (
		user            *dto.User
		errRes          *commondto.ErrorResponse
		response        *commondto.CreateUserResponse
		responseBuilder *builder.Response
		statusCode      int
	)
	defer func() {
		if err := recover(); err != nil {
			responseBuilder := builder.Response{
				Ctx: ctx,
				ErrorRes: &commondto.ErrorResponse{
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

	user, errRes = validator.UserReq(ctx)
	if errRes == nil {
		service := services.CreateUserService(nil)
		response, errRes = service.RegisterUserService(user)
		statusCode = utils.StatusCode[utils.Created]
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
