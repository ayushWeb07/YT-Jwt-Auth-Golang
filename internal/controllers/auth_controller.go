package controllers

import (
	"net/http"

	"github.com/ayushWeb07/YT-Jwt-Auth-Golang/internal/config"
	"github.com/ayushWeb07/YT-Jwt-Auth-Golang/internal/dtos"
	"github.com/ayushWeb07/YT-Jwt-Auth-Golang/internal/services"
	"github.com/ayushWeb07/YT-Jwt-Auth-Golang/internal/utils"
	"go.uber.org/zap"
)

type AuthControllerInterface interface {
	RegisterUser(resWriter http.ResponseWriter, req *http.Request)
}

type AuthController struct {
	AuthService  services.AuthServiceInterface
	logger       *zap.Logger
	serverConfig *config.ServerConfig
}

func (authController *AuthController) RegisterUser(resWriter http.ResponseWriter, req *http.Request) {
	userPayload := req.Context().Value("payload").(*dtos.CreateUserPayload)

	// call the register auth service
	registerUserServiceErr := authController.AuthService.RegisterUser(userPayload)

	if registerUserServiceErr != nil {
		utils.WriteJsonResponse(registerUserServiceErr.StatusCode, resWriter, map[string]any{
			"success": registerUserServiceErr.Success,
			"message": "Something went wrong while user registration",
			"error":   registerUserServiceErr.Error(),
		})

		return
	}

	utils.WriteJsonResponse(http.StatusCreated, resWriter, map[string]any{
		"success": true,
		"message": "Successfully registered the user",
	})
}

func NewAuthController(authService services.AuthServiceInterface, logger *zap.Logger, serverConfig *config.ServerConfig) AuthControllerInterface {

	authController := &AuthController{
		AuthService:  authService,
		logger:       logger,
		serverConfig: serverConfig,
	}

	return authController

}
