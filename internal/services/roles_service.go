package services

import (
	"github.com/ayushWeb07/YT-Jwt-Auth-Golang/internal/config"
	"github.com/ayushWeb07/YT-Jwt-Auth-Golang/internal/repositories"
	"go.uber.org/zap"
)

type RoleServiceInterface interface {
	CreateRole()
	GetAllRoles()
	GetRoleById()
	UpdateRoleById()
	DeleteRoleById()
}

type RoleService struct {
	RoleRepository repositories.RoleRepositoryInterface
	logger         *zap.Logger
	serverConfig   *config.ServerConfig
}

func (roleService *RoleService) CreateRole() {
	roleService.logger.Info("roleService -> CreateRole")

	roleService.RoleRepository.CreateRole()
}

func (roleService *RoleService) GetAllRoles() {
	roleService.logger.Info("roleService -> GetAllRoles")

	roleService.RoleRepository.GetAllRoles()
}

func (roleService *RoleService) GetRoleById() {
	roleService.logger.Info("roleService -> GetRoleById")

	roleService.RoleRepository.GetRoleById()
}

func (roleService *RoleService) UpdateRoleById() {
	roleService.logger.Info("roleService -> UpdateRoleById")

	roleService.RoleRepository.UpdateRoleById()
}

func (roleService *RoleService) DeleteRoleById() {
	roleService.logger.Info("roleService -> DeleteRoleById")

	roleService.RoleRepository.DeleteRoleById()
}

func NewRoleService(RoleRepository repositories.RoleRepositoryInterface, logger *zap.Logger, serverConfig *config.ServerConfig) RoleServiceInterface {

	roleService := &RoleService{
		RoleRepository: RoleRepository,
		logger:         logger,
		serverConfig:   serverConfig,
	}

	return roleService

}
