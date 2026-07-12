package repositories

import (
	"database/sql"

	"github.com/ayushWeb07/YT-Jwt-Auth-Golang/internal/config"
	"go.uber.org/zap"
)

type UserRepositoryInterface interface {
	CreateUser()
	GetAllUsers()
	GetUserById()
	UpdateUserById()
	DeleteUserById()

	GetUserByEmail()
	GetUserByUsernameAndEmail()
}

type UserRepository struct {
	db           *sql.DB
	logger       *zap.Logger
	serverConfig *config.ServerConfig
}

func (userRepository *UserRepository) CreateUser() {
	userRepository.logger.Info("userRepository -> CreateUser")
}

func (userRepository *UserRepository) GetAllUsers() {
	userRepository.logger.Info("userRepository -> GetAllUsers")
}

func (userRepository *UserRepository) GetUserById() {
	userRepository.logger.Info("userRepository -> GetUserById")
}

func (userRepository *UserRepository) UpdateUserById() {
	userRepository.logger.Info("userRepository -> UpdateUserById")
}

func (userRepository *UserRepository) DeleteUserById() {
	userRepository.logger.Info("userRepository -> DeleteUserById")
}

func (userRepository *UserRepository) GetUserByEmail() {
	userRepository.logger.Info("userRepository -> GetUserByEmail")
}

func (userRepository *UserRepository) GetUserByUsernameAndEmail() {
	userRepository.logger.Info("userRepository -> GetUserByUsernameAndEmail")
}

func NewUserRepository(db *sql.DB, logger *zap.Logger, serverConfig *config.ServerConfig) UserRepositoryInterface {

	userRepository := &UserRepository{
		db:           db,
		logger:       logger,
		serverConfig: serverConfig,
	}

	return userRepository

}
