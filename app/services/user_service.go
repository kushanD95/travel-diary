package services

import (
	"fmt"

	"github.com/kushanD95/traval-diary/package/config"
	"github.com/kushanD95/traval-diary/package/dto"
	"github.com/kushanD95/traval-diary/package/utils"
)

func Register(user *dto.User) {
	lg := config.AppConfigutarion.GetLogger()
	// lgFields := []zap.Field{zap.String("Method", "Register")}
	lg.Info(fmt.Sprintf(utils.REGISTER_HANDLER, utils.STARTED))
}
