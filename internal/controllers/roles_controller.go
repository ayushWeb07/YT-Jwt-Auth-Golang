package controllers

import (
	"net/http"

	"github.com/ayushWeb07/YT-Jwt-Auth-Golang/internal/config"
	"github.com/ayushWeb07/YT-Jwt-Auth-Golang/internal/services"
	"github.com/ayushWeb07/YT-Jwt-Auth-Golang/internal/utils"
	"go.uber.org/zap"
)

type RoleControllerInterface interface {
	CreateRole(resWriter http.ResponseWriter, req *http.Request)
	GetAllRoles(resWriter http.ResponseWriter, req *http.Request)
	GetRoleById(resWriter http.ResponseWriter, req *http.Request)
	UpdateRoleById(resWriter http.ResponseWriter, req *http.Request)
	DeleteRoleById(resWriter http.ResponseWriter, req *http.Request)
}

type RoleController struct {
	RoleService  services.RoleServiceInterface
	logger       *zap.Logger
	serverConfig *config.ServerConfig
}

func (roleController *RoleController) CreateRole(resWriter http.ResponseWriter, req *http.Request) {
	roleController.RoleService.CreateRole()

	utils.WriteJsonResponse(http.StatusCreated, resWriter, map[string]any{
		"success": true,
		"message": "A role was created successfully",
	})
}

func (roleController *RoleController) GetAllRoles(resWriter http.ResponseWriter, req *http.Request) {
	roleController.RoleService.GetAllRoles()

	utils.WriteJsonResponse(http.StatusOK, resWriter, map[string]any{
		"success": true,
		"message": "All roles were fetched successfully",
	})
}

func (roleController *RoleController) GetRoleById(resWriter http.ResponseWriter, req *http.Request) {
	roleController.RoleService.GetRoleById()

	utils.WriteJsonResponse(http.StatusOK, resWriter, map[string]any{
		"success": true,
		"message": "A single role was fetched successfully",
	})
}

func (roleController *RoleController) UpdateRoleById(resWriter http.ResponseWriter, req *http.Request) {
	roleController.RoleService.UpdateRoleById()

	utils.WriteJsonResponse(http.StatusOK, resWriter, map[string]any{
		"success": true,
		"message": "A single role was updated successfully",
	})
}

func (roleController *RoleController) DeleteRoleById(resWriter http.ResponseWriter, req *http.Request) {
	roleController.RoleService.DeleteRoleById()

	utils.WriteJsonResponse(http.StatusOK, resWriter, map[string]any{
		"success": true,
		"message": "A single role was deleted successfully",
	})
}

func NewRoleController(RoleService services.RoleServiceInterface, logger *zap.Logger, serverConfig *config.ServerConfig) RoleControllerInterface {

	roleController := &RoleController{
		RoleService:  RoleService,
		logger:       logger,
		serverConfig: serverConfig,
	}

	return roleController

}
