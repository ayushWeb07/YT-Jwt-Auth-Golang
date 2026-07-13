package repositories

import (
	"database/sql"
	"fmt"

	"github.com/ayushWeb07/YT-Jwt-Auth-Golang/internal/config"
	"github.com/ayushWeb07/YT-Jwt-Auth-Golang/internal/database/models"
	"github.com/ayushWeb07/YT-Jwt-Auth-Golang/internal/dtos"
	"github.com/ayushWeb07/YT-Jwt-Auth-Golang/internal/utils"
	"go.uber.org/zap"
)

type UserRepositoryInterface interface {
	CreateUser()
	GetAllUsers() ([]*models.UserModel, *utils.AppError)
	GetUserById(userParams *dtos.GetUserByIdParams) (*models.UserModel, *utils.AppError)
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

func (userRepository *UserRepository) GetAllUsers() ([]*models.UserModel, *utils.AppError) {

	// create an array for the user models
	var userModels []*models.UserModel

	query := "SELECT id, username, email, is_verified, created_at, updated_at FROM users"
	rows, err := userRepository.db.Query(query)

	if err != nil {
		userRepository.logger.Error("Something went wrong while fetching all the users",
			zap.String("error", err.Error()))

		return nil, utils.InternalServerError(fmt.Sprintf("Something went wrong while fetching all the users: %s", err.Error()))
	}

	defer rows.Close()

	for rows.Next() {
		userModel := &models.UserModel{}

		err := rows.Scan(&userModel.ID, &userModel.Username, &userModel.Email, &userModel.IsVerified, &userModel.CreatedAt, &userModel.UpdatedAt)

		if err != nil {
			userRepository.logger.Error("Failed to fetch all users",
				zap.String("error", err.Error()))

			return nil, utils.InternalServerError(fmt.Sprintf("Failed to fetch all users: %s", err.Error()))
		}

		userModels = append(userModels, userModel)
	}

	userRepository.logger.Info("Successfully fetched all the users",
		zap.Int("count", len(userModels)))

	return userModels, nil
}

func (userRepository *UserRepository) GetUserById(userParams *dtos.GetUserByIdParams) (*models.UserModel, *utils.AppError) {

	userModel := &models.UserModel{}

	query := "SELECT id, username, email, is_verified, created_at, updated_at FROM users WHERE id = ?"

	row := userRepository.db.QueryRow(query, userParams.ID)

	err := row.Scan(&userModel.ID, &userModel.Username, &userModel.Email, &userModel.IsVerified, &userModel.CreatedAt, &userModel.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			userRepository.logger.Error("Such user does not exist",
				zap.Int("user_id", userParams.ID))

			return nil, utils.NotFoundError("Such user does not exist")
		}

		userRepository.logger.Error("Failed to fetch the user",
			zap.Int("user_id", userParams.ID),
			zap.String("error", err.Error()))

		return nil, utils.InternalServerError(fmt.Sprintf("Failed to fetch the user: %s", err.Error()))
	}

	userRepository.logger.Info("Successfully fetched the user",
		zap.Int("user_id", userParams.ID))

	return userModel, nil
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
