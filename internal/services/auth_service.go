package services

import (
	"github.com/ayushWeb07/YT-Jwt-Auth-Golang/internal/config"
	"github.com/ayushWeb07/YT-Jwt-Auth-Golang/internal/dtos"
	"github.com/ayushWeb07/YT-Jwt-Auth-Golang/internal/utils"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type AuthServiceInterface interface {
	RegisterUser(userPayload *dtos.CreateUserPayload) *utils.AppError
}

type AuthService struct {
	UserService  UserServiceInterface
	logger       *zap.Logger
	serverConfig *config.ServerConfig
}

func (authService *AuthService) RegisterUser(userPayload *dtos.CreateUserPayload) *utils.AppError {
	authService.logger.Info("authService -> RegisterUser")

	// check if the user already exists
	_, getUserServiceErr := authService.UserService.GetUserByUsernameAndEmail(&dtos.GetUserByUsernameAndEmailPayload{
		Username: userPayload.Username,
		Email:    userPayload.Email,
	})

	if getUserServiceErr == nil {
		return utils.BadRequestError("User with such username and email already exists")
	}

	// hash the password
	hashBytes, hashErr := bcrypt.GenerateFromPassword([]byte(userPayload.Password), bcrypt.DefaultCost)

	if hashErr != nil {
		authService.logger.Fatal("Something went wrong while hashing the password",
			zap.String("error", hashErr.Error()))

		return utils.InternalServerError("Something went wrong while hashing the password: " + hashErr.Error())
	}

	userPayload.Password = string(hashBytes)

	// call create user service
	createUserServiceErr := authService.UserService.CreateUser(userPayload)

	if createUserServiceErr != nil {
		return createUserServiceErr
	}

	authService.logger.Info("Auth registration was successful")
	return nil
}

func NewAuthService(userService UserServiceInterface, logger *zap.Logger, serverConfig *config.ServerConfig) AuthServiceInterface {

	authService := &AuthService{
		UserService:  userService,
		logger:       logger,
		serverConfig: serverConfig,
	}

	return authService

}
