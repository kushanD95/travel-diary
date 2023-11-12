package services

import (
	"fmt"

	"github.com/kushanD95/traval-diary/app/controller/dto"
	"github.com/kushanD95/traval-diary/app/model"
	"github.com/kushanD95/traval-diary/app/repository"
	"github.com/kushanD95/traval-diary/package/config"
	commondto "github.com/kushanD95/traval-diary/package/dto"
	"github.com/kushanD95/traval-diary/package/utils"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserService struct {
	_              struct{}
	serviceContext Context
	transaction    *gorm.DB
}

func CreateUserService(transaction *gorm.DB) UserService {
	userService := UserService{transaction: transaction}
	return userService
}

func (service *UserService) RegisterUserService(user *dto.User) (*commondto.CreateUserResponse, *commondto.ErrorResponse) {
	lg := config.AppConfigutarion.GetLogger()
	lgFields := []zap.Field{zap.String(utils.METHOD, utils.REGISTER_USER_SERVICE)}
	lg.Info(fmt.Sprintf(utils.REGISTER_USER_SERVICE_LOG, utils.STARTED), lgFields...)
	defer func() {
		lg.Info(fmt.Sprintf(utils.REGISTER_USER_SERVICE_LOG, utils.END), lgFields...)
	}()

	userData := &model.User{
		UserName: user.UserName,
		Pwd:      user.Pwd,
	}

	createUserRepo := repository.CreateUserRepository(Db, nil)
	repoUser, err := createUserRepo.RegisterUser(userData)

	if err == nil {
		userId := repoUser.ID
		_ = userId
		// create userdetails used sae transaction & rollback from defer
	}

	if err != nil {
		lgFields = append(lgFields, zap.Any(utils.ERROR, err))
		lg.Error(fmt.Sprintf(utils.REGISTER_USER_SERVICE_LOG, utils.END_WITH_ERROR), lgFields...)
		errRes := service.serviceContext.BuildRepoErrRes(err)
		return nil, errRes
	}

	response := &commondto.CreateUserResponse{
		Status:   0,
		Message:  "",
		Username: "",
	}
	return response, nil
}
