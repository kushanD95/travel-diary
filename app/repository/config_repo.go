package repository

import (
	"github.com/kushanD95/traval-diary/app/model"
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
	envConfig := new(model.EnvConfig)

	err := repo.repoContext.Db.Where("\"ENV_NAME\" = ?", env).First(&envConfig).Error
	if err != nil {
		return nil, err
	}
	return envConfig, nil
}
