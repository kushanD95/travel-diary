package services

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/kushanD95/traval-diary/app/controller/dto"
	"github.com/kushanD95/traval-diary/app/model"
	"github.com/kushanD95/traval-diary/app/repository"
	"github.com/kushanD95/traval-diary/package/config"
	commondto "github.com/kushanD95/traval-diary/package/dto"
	"github.com/kushanD95/traval-diary/package/utils"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
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
	var (
		err             error
		repoUser        *model.User
		repoUserDetails *model.UserDetails
	)
	userData := &model.User{
		UserName:  user.UserName,
		Pwd:       user.Pwd,
		CreatedAt: time.Now().UTC(),
	}

	userDetailsData := &model.UserDetails{
		FName:     user.FName,
		LName:     user.LName,
		Country:   user.Country,
		CreatedAt: time.Now().UTC(),
	}

	repoUser, err = service.CreateUser(userData)

	userDetailsService := CreateUserDetailsService(Db)

	if err == nil {
		userDetailsData.UserID = repoUser.ID
		repoUserDetails, err = userDetailsService.CreateUserDetails(userDetailsData)
		_ = repoUserDetails
		// create userdetails used sae transaction & rollback from defer
	}

	if err != nil {
		lgFields = append(lgFields, zap.Any(utils.ERROR, err))
		lg.Error(fmt.Sprintf(utils.REGISTER_USER_SERVICE_LOG, utils.END_WITH_ERROR), lgFields...)
		errRes := service.serviceContext.BuildRepoErrRes(err)
		return nil, errRes
	}

	response := &commondto.CreateUserResponse{
		Status:   utils.StatusCode[utils.Created],
		Message:  "User Registerd successfuly",
		Username: repoUser.UserName,
	}
	return response, nil
}

func (service *UserService) CreateUser(user *model.User) (*model.User, error) {
	lg := config.AppConfigutarion.GetLogger()
	lgFields := []zap.Field{zap.String(utils.METHOD, utils.CREATE_USER)}
	lg.Info(fmt.Sprintf(utils.CREATE_USER_LOG, utils.STARTED), lgFields...)
	defer func() {
		lg.Info(fmt.Sprintf(utils.CREATE_USER_LOG, utils.END), lgFields...)
	}()

	createUserRepo := repository.CreateUserRepository(Db, nil)
	repoUser, err := createUserRepo.RegisterUser(user)

	return repoUser, err
}

func (service *UserService) LoginUserService(loginUser *commondto.UserLogin) (*commondto.LoginUserResponse, *commondto.ErrorResponse) {
	lg := config.AppConfigutarion.GetLogger()
	lgFields := []zap.Field{zap.String(utils.METHOD, utils.LOGIN_USER_SERVICE)}
	lg.Info(fmt.Sprintf(utils.LOGIN_USER_SERVICE_LOG, utils.STARTED), lgFields...)
	defer func() {
		lg.Info(fmt.Sprintf(utils.LOGIN_USER_SERVICE_LOG, utils.END), lgFields...)
	}()

	createUserRepo := repository.CreateUserRepository(Db, nil)
	repoUser, err := createUserRepo.GetUser(loginUser)

	if err != nil {
		errRes := service.serviceContext.BuildRepoErrRes(err)
		errRes.Message = utils.INVALID_USERNAME
		return nil, errRes
	}

	err = bcrypt.CompareHashAndPassword([]byte(repoUser.Pwd), []byte(loginUser.Pwd))

	if err != nil {
		errRes := &commondto.ErrorResponse{
			Message: utils.INVALID_PWD,
			Code:    utils.StatusCode[utils.BadRequest],
			Error:   fmt.Sprintf("%v", err.Error()),
		}
		return nil, errRes
	}
	token, errRes := service.generateToken(repoUser)

	loginUserResponse := &commondto.LoginUserResponse{
		Status:   utils.StatusCode[utils.Success],
		Message:  utils.SUCCESS_LOGIN,
		Username: repoUser.UserName,
		Token:    token,
	}

	return loginUserResponse, errRes
}

func (service *UserService) generateToken(user *model.User) (string, *commondto.ErrorResponse) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Minute * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(config.AppConfigutarion.Secret))

	if err != nil {
		errRes := &commondto.ErrorResponse{
			Message: utils.FAILED_CREATE_TOKEN,
			Code:    utils.StatusCode[utils.BadRequest],
			Error:   fmt.Sprintf("%v", err.Error()),
		}
		return "", errRes
	}

	return tokenString, nil
}
