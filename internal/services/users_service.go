package services

import (
	"github.com/ayushWeb07/YT-Jwt-Auth-Golang/internal/config"
	"github.com/ayushWeb07/YT-Jwt-Auth-Golang/internal/database/models"
	"github.com/ayushWeb07/YT-Jwt-Auth-Golang/internal/dtos"
	"github.com/ayushWeb07/YT-Jwt-Auth-Golang/internal/repositories"
	"github.com/ayushWeb07/YT-Jwt-Auth-Golang/internal/utils"
	"go.uber.org/zap"
)

type UserServiceInterface interface {
	CreateUser(userPayload *dtos.CreateUserPayload) *utils.AppError
	GetAllUsers() ([]*models.UserModel, *utils.AppError)
	GetUserById(userParams *dtos.GetUserByIdParams) (*models.UserModel, *utils.AppError)
	UpdateUserById()
	DeleteUserById()
	GetUserByUsernameAndEmail(userPayload *dtos.GetUserByUsernameAndEmailPayload) (*models.UserModel, *utils.AppError)
}

type UserService struct {
	UserRepository repositories.UserRepositoryInterface
	logger         *zap.Logger
	serverConfig   *config.ServerConfig
}

func (userService *UserService) CreateUser(userPayload *dtos.CreateUserPayload) *utils.AppError {
	userService.logger.Info("userService -> CreateUser")

	createUserRepositoryErr := userService.UserRepository.CreateUser(userPayload)
	return createUserRepositoryErr
}

func (userService *UserService) GetAllUsers() ([]*models.UserModel, *utils.AppError) {
	userService.logger.Info("userService -> GetAllUsers")

	userModels, err := userService.UserRepository.GetAllUsers()
	return userModels, err
}

func (userService *UserService) GetUserById(userParams *dtos.GetUserByIdParams) (*models.UserModel, *utils.AppError) {
	userService.logger.Info("userService -> GetUserById")

	userModel, err := userService.UserRepository.GetUserById(userParams)
	return userModel, err
}

func (userService *UserService) UpdateUserById() {
	userService.logger.Info("userService -> UpdateUserById")

	userService.UserRepository.UpdateUserById()
}

func (userService *UserService) DeleteUserById() {
	userService.logger.Info("userService -> DeleteUserById")

	userService.UserRepository.DeleteUserById()
}

func (userService *UserService) GetUserByUsernameAndEmail(userPayload *dtos.GetUserByUsernameAndEmailPayload) (*models.UserModel, *utils.AppError) {
	userService.logger.Info("userService -> GetUserByUsernameAndEmail")

	userModel, err := userService.UserRepository.GetUserByUsernameAndEmail(userPayload)
	return userModel, err
}

func NewUserService(UserRepository repositories.UserRepositoryInterface, logger *zap.Logger, serverConfig *config.ServerConfig) UserServiceInterface {

	userService := &UserService{
		UserRepository: UserRepository,
		logger:         logger,
		serverConfig:   serverConfig,
	}

	return userService

}
