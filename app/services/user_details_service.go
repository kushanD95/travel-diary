package services

import (
	"fmt"

	"github.com/kushanD95/traval-diary/app/model"
	"github.com/kushanD95/traval-diary/app/repository"
	"github.com/kushanD95/traval-diary/package/config"
	"github.com/kushanD95/traval-diary/package/utils"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserDetailsService struct {
	_              struct{}
	serviceContext Context
	transaction    *gorm.DB
}

func CreateUserDetailsService(transaction *gorm.DB) UserDetailsService {
	userDetailsService := UserDetailsService{transaction: transaction}
	return userDetailsService
}

func (service *UserDetailsService) CreateUserDetails(userDetails *model.UserDetails) (*model.UserDetails, error) {
	lg := config.AppConfigutarion.GetLogger()
	lgFields := []zap.Field{zap.String(utils.METHOD, utils.CREATE_USER)}
	lg.Info(fmt.Sprintf(utils.CREATE_USER_LOG, utils.STARTED), lgFields...)
	defer func() {
		lg.Info(fmt.Sprintf(utils.CREATE_USER_LOG, utils.END), lgFields...)
	}()

	createUserDetailsRepo := repository.CreateUserDetailsRepository(Db, Db)
	repoUser, err := createUserDetailsRepo.CreateUserDetails(userDetails)

	return repoUser, err
}
