package validator

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/kushanD95/traval-diary/app/controller/dto"
	commonDto "github.com/kushanD95/traval-diary/package/dto"
	"github.com/kushanD95/traval-diary/package/utils"
)

func EnvConfigReq(ctx *fiber.Ctx) (*dto.EnvConfigReq, *commonDto.ErrorResponse) {
	var (
		request dto.EnvConfigReq
	)
	parseErr := ctx.BodyParser(&request)
	if parseErr != nil {
		errRes := &commonDto.ErrorResponse{
			Message: "Invalid Request body",
			Code:    utils.StatusCode[utils.BadRequest],
			Error:   fmt.Sprintf("%v", parseErr.Error()),
		}
		return nil, errRes
	}
	validateErr := validate.Struct(&request)
	if validateErr != nil {
		errRes := &commonDto.ErrorResponse{
			Message: "Invalid Request body",
			Code:    utils.StatusCode[utils.BadRequest],
			Error:   fmt.Sprintf("%v", validateErr.Error()),
		}
		return nil, errRes
	}

	return &request, nil
}
