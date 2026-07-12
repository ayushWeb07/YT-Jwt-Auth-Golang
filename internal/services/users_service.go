package services

import (
	"github.com/ayushWeb07/YT-Jwt-Auth-Golang/internal/config"
	"github.com/ayushWeb07/YT-Jwt-Auth-Golang/internal/repositories"
	"go.uber.org/zap"
)

type UserServiceInterface interface {
	CreateUser()
	GetAllUsers()
	GetUserById()
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

func (userService *UserService) GetAllUsers() {
	userService.logger.Info("userService -> GetAllUsers")

	userService.UserRepository.GetAllUsers()
}

func (userService *UserService) GetUserById() {
	userService.logger.Info("userService -> GetUserById")

	userService.UserRepository.GetUserById()
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
