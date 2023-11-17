package validator

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/kushanD95/traval-diary/app/controller/dto"
	"github.com/kushanD95/traval-diary/package/config"
	commondto "github.com/kushanD95/traval-diary/package/dto"
	"github.com/kushanD95/traval-diary/package/utils"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

func genHash(pwd string) (string, *commondto.ErrorResponse) {
	lg := config.AppConfigutarion.GetLogger()
	lgFields := []zap.Field{zap.String(utils.METHOD, utils.GEN_HASH)}
	lg.Info(fmt.Sprintf(utils.GEN_HASH_LOG, utils.STARTED), lgFields...)

	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), 10)

	if err != nil {
		errRes := &commondto.ErrorResponse{
			Message: utils.INVALID_REQUEST_BODY,
			Code:    utils.StatusCode[utils.BadRequest],
			Error:   fmt.Sprintf("%v", err.Error()),
		}
		lgFields = append(lgFields, zap.Any(utils.ERROR, err))
		lg.Error(fmt.Sprintf(utils.GEN_HASH_LOG, utils.END_WITH_ERROR), lgFields...)
		return "", errRes
	}

	lg.Info(fmt.Sprintf(utils.GEN_HASH_LOG, utils.END), lgFields...)
	return string(hash), nil
}

func UserReq(ctx *fiber.Ctx) (*dto.User, *commondto.ErrorResponse) {
	lg := config.AppConfigutarion.GetLogger()
	lgFields := []zap.Field{zap.String(utils.METHOD, utils.USER_REQ)}
	lg.Info(fmt.Sprintf(utils.USER_REQ_LOG, utils.STARTED), lgFields...)
	var (
		request dto.User
	)
	parseErr := ctx.BodyParser(&request)
	if parseErr != nil {
		errRes := &commondto.ErrorResponse{
			Message: utils.INVALID_REQUEST_BODY,
			Code:    utils.StatusCode[utils.BadRequest],
			Error:   fmt.Sprintf("%v", parseErr.Error()),
		}

		lgFields = append(lgFields, zap.Any(utils.ERROR, parseErr))
		lg.Error(fmt.Sprintf(utils.USER_REQ_LOG, utils.END_WITH_ERROR), lgFields...)
		return nil, errRes
	}
	validateErr := validate.Struct(&request)
	if validateErr != nil {
		errRes := &commondto.ErrorResponse{
			Message: utils.INVALID_REQUEST_BODY,
			Code:    utils.StatusCode[utils.BadRequest],
			Error:   fmt.Sprintf("%v", validateErr.Error()),
		}
		lgFields = append(lgFields, zap.Any(utils.ERROR, validateErr))
		lg.Error(fmt.Sprintf(utils.USER_REQ_LOG, utils.END_WITH_ERROR), lgFields...)
		return nil, errRes
	}

	hash, errRes := genHash(request.Pwd)

	if errRes != nil {
		lgFields = append(lgFields, zap.Any(utils.ERROR, errRes.Error))
		lg.Error(fmt.Sprintf(utils.USER_REQ_LOG, utils.END_WITH_ERROR), lgFields...)
		return nil, errRes
	}

	request.Pwd = hash

	lg.Info(fmt.Sprintf(utils.USER_REQ_LOG, utils.END), lgFields...)
	return &request, nil
}

func UserLoginReq(ctx *fiber.Ctx) (*commondto.UserLogin, *commondto.ErrorResponse) {
	lg := config.AppConfigutarion.GetLogger()
	lgFields := []zap.Field{zap.String(utils.METHOD, utils.USER_LOGIN_REQ)}
	lg.Info(fmt.Sprintf(utils.USER_LOGIN_REQ_LOG, utils.STARTED), lgFields...)
	var (
		request          dto.UserLogin
		validatedRequest *commondto.UserLogin
	)
	parseErr := ctx.BodyParser(&request)
	if parseErr != nil {
		errRes := &commondto.ErrorResponse{
			Message: utils.INVALID_REQUEST_BODY,
			Code:    utils.StatusCode[utils.BadRequest],
			Error:   fmt.Sprintf("%v", parseErr.Error()),
		}

		lgFields = append(lgFields, zap.Any(utils.ERROR, parseErr))
		lg.Error(fmt.Sprintf(utils.USER_LOGIN_REQ_LOG, utils.END_WITH_ERROR), lgFields...)
		return nil, errRes
	}
	validateErr := validate.Struct(&request)
	if validateErr != nil {
		errRes := &commondto.ErrorResponse{
			Message: utils.INVALID_REQUEST_BODY,
			Code:    utils.StatusCode[utils.BadRequest],
			Error:   fmt.Sprintf("%v", validateErr.Error()),
		}
		lgFields = append(lgFields, zap.Any(utils.ERROR, validateErr))
		lg.Error(fmt.Sprintf(utils.USER_LOGIN_REQ_LOG, utils.END_WITH_ERROR), lgFields...)
		return nil, errRes
	}

	validatedRequest = &commondto.UserLogin{
		UserName: request.UserName,
		Pwd:      request.Pwd,
	}
	lg.Info(fmt.Sprintf(utils.USER_LOGIN_REQ_LOG, utils.END), lgFields...)
	return validatedRequest, nil
}
