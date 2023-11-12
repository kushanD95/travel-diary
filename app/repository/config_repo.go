package repository

import (
	"fmt"

	"github.com/kushanD95/traval-diary/app/model"
	"github.com/kushanD95/traval-diary/package/config"
	"github.com/kushanD95/traval-diary/package/utils"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ConfigRepository struct {
	_           struct{}
	repoContext Context
}

func CreateConfigRepository(db *gorm.DB, transaction *gorm.DB) ConfigRepository {
	repoContext := CreateRepositoryContext(db, transaction)
	configRepo := ConfigRepository{repoContext: repoContext}
	return configRepo
}

func (repo *ConfigRepository) GetConfig(env string) (*model.EnvConfig, error) {
	lg := config.AppConfigutarion.GetLogger()
	lgFields := []zap.Field{zap.String(utils.METHOD, utils.GET_CONFIG)}
	lg.Info(fmt.Sprintf(utils.GET_CONFIG_REPO, utils.STARTED), lgFields...)
	envConfig := new(model.EnvConfig)

	err := repo.repoContext.Db.Where(ENV_NAME, env).First(&envConfig).Error
	if err != nil {
		lgFields = append(lgFields, zap.Any(utils.ERROR, err))
		lg.Error(fmt.Sprintf(utils.GET_CONFIG_REPO, utils.END_WITH_ERROR), lgFields...)
		return nil, err
	}
	lg.Info(fmt.Sprintf(utils.GET_CONFIG_REPO, utils.END), lgFields...)
	return envConfig, nil
}
