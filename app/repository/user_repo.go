package repository

import (
	"fmt"

	"github.com/kushanD95/traval-diary/app/model"
	"github.com/kushanD95/traval-diary/package/config"
	commondto "github.com/kushanD95/traval-diary/package/dto"
	"github.com/kushanD95/traval-diary/package/utils"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserRepository struct {
	_           struct{}
	repoContext Context
}

func CreateUserRepository(db *gorm.DB, transaction *gorm.DB) UserRepository {
	repoContext := CreateRepositoryContext(db, transaction)
	userRepo := UserRepository{repoContext: repoContext}
	return userRepo
}

func (repo *UserRepository) RegisterUser(user *model.User) (*model.User, error) {
	lg := config.AppConfigutarion.GetLogger()
	lgFields := []zap.Field{zap.String(utils.METHOD, utils.REGISTER_USER)}
	lg.Info(fmt.Sprintf(utils.REGISTER_USER_REPO, utils.STARTED), lgFields...)
	result := repo.repoContext.Db.Table(user.TableName()).Create(&user)
	err := result.Error
	if err != nil {
		lgFields = append(lgFields, zap.Any(utils.ERROR, err))
		lg.Error(fmt.Sprintf(utils.REGISTER_USER_REPO, utils.END_WITH_ERROR), lgFields...)
		return nil, err
	}

	lg.Info(fmt.Sprintf(utils.REGISTER_USER_REPO, utils.END), lgFields...)
	return user, nil
}

func (repo *UserRepository) GetUser(userRequest *commondto.UserLogin) (*model.User, error) {
	lg := config.AppConfigutarion.GetLogger()
	lgFields := []zap.Field{zap.String(utils.METHOD, utils.GET_USER)}
	lg.Info(fmt.Sprintf(utils.GET_USER_REPO, utils.STARTED), lgFields...)
	var user *model.User
	result := repo.repoContext.Db.First(&user, USER_NAME, userRequest.UserName)
	err := result.Error
	if err != nil {
		lgFields = append(lgFields, zap.Any(utils.ERROR, err))
		lg.Error(fmt.Sprintf(utils.GET_USER_REPO, utils.END_WITH_ERROR), lgFields...)
		return nil, err
	}

	lg.Info(fmt.Sprintf(utils.GET_USER_REPO, utils.END), lgFields...)
	return user, nil
}
