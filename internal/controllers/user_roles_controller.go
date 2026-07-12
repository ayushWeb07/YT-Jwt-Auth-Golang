package controllers

import (
	"net/http"

	"github.com/ayushWeb07/YT-Jwt-Auth-Golang/internal/config"
	"github.com/ayushWeb07/YT-Jwt-Auth-Golang/internal/services"
	"github.com/ayushWeb07/YT-Jwt-Auth-Golang/internal/utils"
	"go.uber.org/zap"
)

type UserRoleControllerInterface interface {
	GetRolesOfUser(resWriter http.ResponseWriter, req *http.Request)
	AssignRoleToUser(resWriter http.ResponseWriter, req *http.Request)
	RemoveRoleFromUser(resWriter http.ResponseWriter, req *http.Request)
	CheckUserHasSingleRole(resWriter http.ResponseWriter, req *http.Request)
	CheckUserHasAnyRoles(resWriter http.ResponseWriter, req *http.Request)
	CheckUserHasAllRoles(resWriter http.ResponseWriter, req *http.Request)
}

type UserRoleController struct {
	UserRoleService services.UserRoleServiceInterface
	logger          *zap.Logger
	serverConfig    *config.ServerConfig
}

func (userRoleController *UserRoleController) GetRolesOfUser(resWriter http.ResponseWriter, req *http.Request) {
	userRoleController.UserRoleService.GetRolesOfUser()

	utils.WriteJsonResponse(http.StatusOK, resWriter, map[string]any{
		"success": true,
		"message": "Fetched roles of the user",
	})
}

func (userRoleController *UserRoleController) AssignRoleToUser(resWriter http.ResponseWriter, req *http.Request) {
	userRoleController.UserRoleService.AssignRoleToUser()

	utils.WriteJsonResponse(http.StatusOK, resWriter, map[string]any{
		"success": true,
		"message": "Assign role to the user",
	})
}

func (userRoleController *UserRoleController) RemoveRoleFromUser(resWriter http.ResponseWriter, req *http.Request) {
	userRoleController.UserRoleService.RemoveRoleFromUser()

	utils.WriteJsonResponse(http.StatusOK, resWriter, map[string]any{
		"success": true,
		"message": "Removed role from the user",
	})
}

func (userRoleController *UserRoleController) CheckUserHasSingleRole(resWriter http.ResponseWriter, req *http.Request) {
	userRoleController.UserRoleService.CheckUserHasSingleRole()

	utils.WriteJsonResponse(http.StatusOK, resWriter, map[string]any{
		"success": true,
		"message": "Check user has single role",
	})
}

func (userRoleController *UserRoleController) CheckUserHasAnyRoles(resWriter http.ResponseWriter, req *http.Request) {
	userRoleController.UserRoleService.CheckUserHasAnyRoles()

	utils.WriteJsonResponse(http.StatusOK, resWriter, map[string]any{
		"success": true,
		"message": "Check user has any roles",
	})
}

func (userRoleController *UserRoleController) CheckUserHasAllRoles(resWriter http.ResponseWriter, req *http.Request) {
	userRoleController.UserRoleService.CheckUserHasAllRoles()

	utils.WriteJsonResponse(http.StatusOK, resWriter, map[string]any{
		"success": true,
		"message": "Check user has all roles",
	})
}

func NewUserRoleController(UserRoleService services.UserRoleServiceInterface, logger *zap.Logger, serverConfig *config.ServerConfig) UserRoleControllerInterface {

	userRoleController := &UserRoleController{
		UserRoleService: UserRoleService,
		logger:          logger,
		serverConfig:    serverConfig,
	}

	return userRoleController

}
