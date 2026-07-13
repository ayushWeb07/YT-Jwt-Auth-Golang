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
	CreateUser()
	GetAllUsers() ([]*models.UserModel, *utils.AppError)
	GetUserById(userParams *dtos.GetUserByIdParams) (*models.UserModel, *utils.AppError)
	UpdateUserById()
	DeleteUserById()

	GetUserByEmail()
	GetUserByUsernameAndEmail()
}

type UserService struct {
	UserRepository repositories.UserRepositoryInterface
	logger         *zap.Logger
	serverConfig   *config.ServerConfig
}

func (userService *UserService) CreateUser() {
	userService.logger.Info("userService -> CreateUser")

	userService.UserRepository.CreateUser()
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

func (userService *UserService) GetUserByEmail() {
	userService.logger.Info("userService -> GetUserByEmail")

	userService.UserRepository.GetUserByEmail()
}

func (userService *UserService) GetUserByUsernameAndEmail() {
	userService.logger.Info("userService -> GetUserByUsernameAndEmail")

	userService.UserRepository.GetUserByUsernameAndEmail()
}

func NewUserService(UserRepository repositories.UserRepositoryInterface, logger *zap.Logger, serverConfig *config.ServerConfig) UserServiceInterface {

	userService := &UserService{
		UserRepository: UserRepository,
		logger:         logger,
		serverConfig:   serverConfig,
	}

	return userService

}
