package services

import (
	"github.com/ayushWeb07/YT-Jwt-Auth-Golang/internal/config"
	"github.com/ayushWeb07/YT-Jwt-Auth-Golang/internal/database/models"
	"github.com/ayushWeb07/YT-Jwt-Auth-Golang/internal/dtos"
	"github.com/ayushWeb07/YT-Jwt-Auth-Golang/internal/repositories"
	"github.com/ayushWeb07/YT-Jwt-Auth-Golang/internal/utils"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceInterface interface {
	CreateUser(userPayload *dtos.CreateUserPayload) *utils.AppError
	GetAllUsers() ([]*models.UserModel, *utils.AppError)
	GetUserById(userParams *dtos.GetUserByIdParams) (*models.UserModel, *utils.AppError)
	UpdateUserById()
	DeleteUserById()
}

type UserService struct {
	UserRepository repositories.UserRepositoryInterface
	logger         *zap.Logger
	serverConfig   *config.ServerConfig
}

func (userService *UserService) CreateUser(userPayload *dtos.CreateUserPayload) *utils.AppError {
	userService.logger.Info("userService -> CreateUser")

	// check if the user already exists
	_, getUserRepositoryErr := userService.UserRepository.GetUserByUsernameAndEmail(&dtos.GetUserByUsernameAndEmailPayload{
		Username: userPayload.Username,
		Email:    userPayload.Email,
	})

	if getUserRepositoryErr == nil {
		return utils.BadRequestError("User with such username and email already exists")
	}

	// hash the password
	hashBytes, hashErr := bcrypt.GenerateFromPassword([]byte(userPayload.Password), bcrypt.DefaultCost)

	if hashErr != nil {
		userService.logger.Fatal("Something went wrong while hashing the password",
			zap.String("error", hashErr.Error()))

		return utils.InternalServerError("Something went wrong while hashing the password: " + hashErr.Error())
	}

	userPayload.Password = string(hashBytes)

	// call create user repository
	createUserRepositoryErr := userService.UserRepository.CreateUser(userPayload)

	if createUserRepositoryErr != nil {
		return createUserRepositoryErr
	}

	userService.logger.Info("User creation was successful")
	return nil
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

func NewUserService(UserRepository repositories.UserRepositoryInterface, logger *zap.Logger, serverConfig *config.ServerConfig) UserServiceInterface {

	userService := &UserService{
		UserRepository: UserRepository,
		logger:         logger,
		serverConfig:   serverConfig,
	}

	return userService

}
