package services

import (
	"fmt"

	"github.com/kushanD95/traval-diary/app/repository"

	"github.com/kushanD95/traval-diary/app/controller/dto"
	"github.com/kushanD95/traval-diary/package/config"
	commondto "github.com/kushanD95/traval-diary/package/dto"
	"github.com/kushanD95/traval-diary/package/utils"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ConfigService struct {
	_              struct{}
	serviceContext Context
	transaction    *gorm.DB
}

func CreateConfigService(transaction *gorm.DB) ConfigService {
	configService := ConfigService{transaction: transaction}
	return configService
}

func (service *ConfigService) FetchConfigService(req *dto.EnvConfigReq) (*commondto.EnvConfig, *commondto.ErrorResponse) {

	lg := config.AppConfigutarion.GetLogger()
	lgFields := []zap.Field{zap.String("Method", "FetchConfigService")}
	lg.Info(fmt.Sprintf(utils.FETCH_CONFIG_SERVICE, utils.STARTED))

	defer func() {
		lg.Info(fmt.Sprintf(utils.FETCH_CONFIG_SERVICE, utils.END), lgFields...)
	}()

	repo := repository.CreateConfigRepository(Db, nil)

	data, err := repo.GetConfig(req.Env)
	if err != nil {
		errRes := service.serviceContext.BuildRepoErrRes(err)
		return nil, errRes
	}
	response := &commondto.EnvConfig{
		EnvURL:  data.EnvURL,
		EnvName: data.EnvName,
	}
	return response, nil
}
