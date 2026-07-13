package controllers

import (
	"net/http"

	"github.com/ayushWeb07/YT-Jwt-Auth-Golang/internal/config"
	"github.com/ayushWeb07/YT-Jwt-Auth-Golang/internal/dtos"
	"github.com/ayushWeb07/YT-Jwt-Auth-Golang/internal/services"
	"github.com/ayushWeb07/YT-Jwt-Auth-Golang/internal/utils"
	"go.uber.org/zap"
)

type UserControllerInterface interface {
	CreateUser(resWriter http.ResponseWriter, req *http.Request)
	GetAllUsers(resWriter http.ResponseWriter, req *http.Request)
	GetUserById(resWriter http.ResponseWriter, req *http.Request)
	UpdateUserById(resWriter http.ResponseWriter, req *http.Request)
	DeleteUserById(resWriter http.ResponseWriter, req *http.Request)

	GetUserByEmail(resWriter http.ResponseWriter, req *http.Request)
	GetUserByUsernameAndEmail(resWriter http.ResponseWriter, req *http.Request)
}

type UserController struct {
	UserService  services.UserServiceInterface
	logger       *zap.Logger
	serverConfig *config.ServerConfig
}

func (userController *UserController) CreateUser(resWriter http.ResponseWriter, req *http.Request) {
	userController.UserService.CreateUser()

	utils.WriteJsonResponse(http.StatusCreated, resWriter, map[string]any{
		"success": true,
		"message": "A user was created successfully",
	})
}

func (userController *UserController) GetAllUsers(resWriter http.ResponseWriter, req *http.Request) {
	userModels, err := userController.UserService.GetAllUsers()

	if err != nil {
		utils.WriteJsonResponse(err.StatusCode, resWriter, map[string]any{
			"success": err.Success,
			"message": err.Error(),
		})

		return
	}

	utils.WriteJsonResponse(http.StatusOK, resWriter, map[string]any{
		"success": true,
		"message": "Successfully fetched all the users",
		"data":    userModels,
	})
}

func (userController *UserController) GetUserById(resWriter http.ResponseWriter, req *http.Request) {
	userParams := req.Context().Value("params").(*dtos.GetUserByIdParams)

	userModel, err := userController.UserService.GetUserById(userParams)

	if err != nil {
		utils.WriteJsonResponse(err.StatusCode, resWriter, map[string]any{
			"success": err.Success,
			"message": err.Error(),
		})

		return
	}

	utils.WriteJsonResponse(http.StatusOK, resWriter, map[string]any{
		"success": true,
		"message": "Successfully fetched the user",
		"data":    userModel,
	})
}

func (userController *UserController) UpdateUserById(resWriter http.ResponseWriter, req *http.Request) {
	userController.UserService.UpdateUserById()

	utils.WriteJsonResponse(http.StatusOK, resWriter, map[string]any{
		"success": true,
		"message": "A user was updated successfully",
	})
}

func (userController *UserController) DeleteUserById(resWriter http.ResponseWriter, req *http.Request) {
	userController.UserService.DeleteUserById()

	utils.WriteJsonResponse(http.StatusOK, resWriter, map[string]any{
		"success": true,
		"message": "A user was deleted successfully",
	})
}

func (userController *UserController) GetUserByEmail(resWriter http.ResponseWriter, req *http.Request) {
	userController.UserService.GetUserByEmail()

	utils.WriteJsonResponse(http.StatusOK, resWriter, map[string]any{
		"success": true,
		"message": "Fetched user successfully by email",
	})
}

func (userController *UserController) GetUserByUsernameAndEmail(resWriter http.ResponseWriter, req *http.Request) {
	userController.UserService.GetUserByUsernameAndEmail()

	utils.WriteJsonResponse(http.StatusOK, resWriter, map[string]any{
		"success": true,
		"message": "Fetched user successfully by username and email",
	})
}

func NewUserController(UserService services.UserServiceInterface, logger *zap.Logger, serverConfig *config.ServerConfig) UserControllerInterface {

	userController := &UserController{
		UserService:  UserService,
		logger:       logger,
		serverConfig: serverConfig,
	}

	return userController

}
