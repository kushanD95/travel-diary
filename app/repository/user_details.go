package repository

import (
	"fmt"

	"github.com/kushanD95/traval-diary/app/model"
	"github.com/kushanD95/traval-diary/package/config"
	"github.com/kushanD95/traval-diary/package/utils"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserDetailsRepository struct {
	_           struct{}
	repoContext Context
}

func CreateUserDetailsRepository(db *gorm.DB, transaction *gorm.DB) UserDetailsRepository {
	repoContext := CreateRepositoryContext(db, transaction)
	userDetailsRepo := UserDetailsRepository{repoContext: repoContext}
	return userDetailsRepo
}

func (repo *UserDetailsRepository) CreateUserDetails(userDetails *model.UserDetails) (*model.UserDetails, error) {
	lg := config.AppConfigutarion.GetLogger()
	lgFields := []zap.Field{zap.String(utils.METHOD, utils.CREATE_USER_DETAILS)}
	lg.Info(fmt.Sprintf(utils.CREATE_USER_DETAILS_REPO, utils.STARTED), lgFields...)
	repoUserDetails := new(model.UserDetails)
	err := repo.repoContext.Db.Table(userDetails.TableName()).Create(&userDetails).Error
	if err != nil {
		lgFields = append(lgFields, zap.Any(utils.ERROR, err))
		lg.Error(fmt.Sprintf(utils.CREATE_USER_DETAILS_REPO, utils.END_WITH_ERROR), lgFields...)
		return nil, err
	}
	lg.Info(fmt.Sprintf(utils.CREATE_USER_DETAILS_REPO, utils.END), lgFields...)
	return repoUserDetails, nil
}
