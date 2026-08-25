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
	CreateUser(userPayload *dtos.CreateUserPayload) *utils.AppError
	GetAllUsers() ([]*models.UserModel, *utils.AppError)
	GetUserById(userParams *dtos.GetUserByIdParams) (*models.UserModel, *utils.AppError)
	UpdateUserById()
	DeleteUserById()

	GetUserByEmail(userPayload *dtos.GetUserByEmailPayload) (*models.UserModel, *utils.AppError)
	GetUserByUsernameAndEmail(userPayload *dtos.GetUserByUsernameAndEmailPayload) (*models.UserModel, *utils.AppError)
}

type UserRepository struct {
	db           *sql.DB
	logger       *zap.Logger
	serverConfig *config.ServerConfig
}

func (userRepository *UserRepository) CreateUser(userPayload *dtos.CreateUserPayload) *utils.AppError {
	query := "INSERT INTO users (username, email, password) VALUES (?, ?, ?)"
	result, queryExecErr := userRepository.db.Exec(query, userPayload.Username, userPayload.Email, userPayload.Password)

	if queryExecErr != nil {
		userRepository.logger.Error("Failed to insert user into the database",
			zap.String("error", queryExecErr.Error()))

		return utils.InternalServerError("Failed to insert user into the database: " + queryExecErr.Error())
	}

	lastInsertedId, insertErr := result.LastInsertId()

	if insertErr != nil {
		userRepository.logger.Error("Failed to insert user into the database",
			zap.String("error", insertErr.Error()))

		return utils.InternalServerError("Failed to insert user into the database: " + insertErr.Error())
	}

	userRepository.logger.Info("Successfully inserted user into the database",
		zap.Int64("user_id", lastInsertedId))

	return nil
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

func (userRepository *UserRepository) GetUserByEmail(userPayload *dtos.GetUserByEmailPayload) (*models.UserModel, *utils.AppError) {
	userModel := &models.UserModel{}

	query := "SELECT id, username, email, is_verified, created_at, updated_at FROM users WHERE email = ?"

	row := userRepository.db.QueryRow(query, userPayload.Email)

	queryErr := row.Scan(&userModel.ID, &userModel.Username, &userModel.Email, &userModel.IsVerified, &userModel.CreatedAt, &userModel.UpdatedAt)

	if queryErr != nil {
		if queryErr == sql.ErrNoRows {
			userRepository.logger.Error("Such user does not exist",
				zap.String("error", queryErr.Error()))

			return nil, utils.NotFoundError("Such user does not exist")
		}

		userRepository.logger.Error("Failed to fetch the user",
			zap.String("error", queryErr.Error()))

		return nil, utils.InternalServerError(fmt.Sprintf("Failed to fetch the user: %s", queryErr.Error()))
	}

	return userModel, nil
}

func (userRepository *UserRepository) GetUserByUsernameAndEmail(userPayload *dtos.GetUserByUsernameAndEmailPayload) (*models.UserModel, *utils.AppError) {
	userModel := &models.UserModel{}

	query := "SELECT id, username, email, is_verified, created_at, updated_at FROM users WHERE username = ? AND email = ?"

	row := userRepository.db.QueryRow(query, userPayload.Username, userPayload.Email)

	queryErr := row.Scan(&userModel.ID, &userModel.Username, &userModel.Email, &userModel.IsVerified, &userModel.CreatedAt, &userModel.UpdatedAt)

	if queryErr != nil {
		if queryErr == sql.ErrNoRows {
			userRepository.logger.Error("Such user does not exist",
				zap.String("error", queryErr.Error()))

			return nil, utils.NotFoundError("Such user does not exist")
		}

		userRepository.logger.Error("Failed to fetch the user",
			zap.String("error", queryErr.Error()))

		return nil, utils.InternalServerError(fmt.Sprintf("Failed to fetch the user: %s", queryErr.Error()))
	}

	return userModel, nil
}

func NewUserRepository(db *sql.DB, logger *zap.Logger, serverConfig *config.ServerConfig) UserRepositoryInterface {

	userRepository := &UserRepository{
		db:           db,
		logger:       logger,
		serverConfig: serverConfig,
	}

	return userRepository

}
