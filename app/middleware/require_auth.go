package middleware

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/kushanD95/traval-diary/app/model"
	"github.com/kushanD95/traval-diary/app/services"
	"github.com/kushanD95/traval-diary/package/config"
	commondto "github.com/kushanD95/traval-diary/package/dto"
	"github.com/kushanD95/traval-diary/package/utils"
	"go.uber.org/zap"
)

func RequireAuth(ctx *fiber.Ctx) error {
	lg := config.AppConfigutarion.GetLogger()
	lgFields := []zap.Field{zap.String(utils.METHOD, utils.REQUIRE_AUTH)}
	lg.Info(fmt.Sprintf(utils.REQUIRE_AUTH_LOG, utils.STARTED), lgFields...)
	var tokenString string

	authorization := ctx.Get(AUTHORIZATION)

	if len(authorization) > 7 && strings.HasPrefix(authorization, BEARER_SPACE) {
		tokenString = strings.TrimPrefix(authorization, BEARER_SPACE)
	} else {
		lg.Warn(fmt.Sprintf(utils.REQUIRE_AUTH_LOG, utils.END_WITH_ERROR), lgFields...)
		errRes := &commondto.ErrorResponse{
			Message: utils.INVALID_BEARER_TOKEN,
			Code:    utils.StatusCode[utils.Unauthorized],
			Error:   UNAUTHORIZED,
		}
		return ctx.Status(fiber.StatusUnauthorized).JSON(errRes)
	}

	tokenByte, err := jwt.Parse(tokenString, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			lg.Warn(fmt.Sprintf(utils.REQUIRE_AUTH_LOG, utils.END_WITH_ERROR), lgFields...)
			return nil, fmt.Errorf(utils.UNEXPECTED_SIGNING_METHOD, jwtToken.Header[ALG])
		}

		return []byte(config.AppConfigutarion.Secret), nil
	})
	if err != nil {
		lg.Warn(fmt.Sprintf(utils.REQUIRE_AUTH_LOG, utils.END_WITH_ERROR), lgFields...)
		errRes := &commondto.ErrorResponse{
			Message: utils.INVALID_BEARER_TOKEN,
			Code:    utils.StatusCode[utils.Unauthorized],
			Error:   err.Error(),
		}
		return ctx.Status(fiber.StatusUnauthorized).JSON(errRes)
	}

	claims, ok := tokenByte.Claims.(jwt.MapClaims)
	if !ok || !tokenByte.Valid {
		lg.Warn(fmt.Sprintf(utils.REQUIRE_AUTH_LOG, utils.END_WITH_ERROR), lgFields...)
		errRes := &commondto.ErrorResponse{
			Message: utils.INVALID_TOKEN_CLAIM,
			Code:    utils.StatusCode[utils.Unauthorized],
			Error:   UNAUTHORIZED,
		}
		return ctx.Status(fiber.StatusUnauthorized).JSON(errRes)

	}

	var user model.User
	services.Db.First(&user, "\"ID\" = ?", fmt.Sprint(claims[SUB])) //TODO [handler db error]

	if float64(user.ID) != claims[SUB] {
		lg.Warn(fmt.Sprintf(utils.REQUIRE_AUTH_LOG, utils.END_WITH_ERROR), lgFields...)
		errRes := &commondto.ErrorResponse{
			Message: utils.INVALID_TOKEN_CLAIM,
			Code:    utils.StatusCode[utils.Forbidden],
			Error:   UNAUTHORIZED,
		}
		return ctx.Status(fiber.StatusForbidden).JSON(errRes)
	}

	lg.Info(fmt.Sprintf(utils.REQUIRE_AUTH_LOG, utils.END), lgFields...)
	return ctx.Next()
}
