package services

import (
	"github.com/ayushWeb07/YT-Jwt-Auth-Golang/internal/config"
	"github.com/ayushWeb07/YT-Jwt-Auth-Golang/internal/repositories"
	"go.uber.org/zap"
)

type UserRoleServiceInterface interface {
	GetRolesOfUser()
	AssignRoleToUser()
	RemoveRoleFromUser()
	CheckUserHasSingleRole()
	CheckUserHasAnyRoles()
	CheckUserHasAllRoles()
}

type UserRoleService struct {
	UserRoleRepository repositories.UserRoleRepositoryInterface
	logger             *zap.Logger
	serverConfig       *config.ServerConfig
}

func (userRoleService *UserRoleService) GetRolesOfUser() {
	userRoleService.logger.Info("userRoleService -> GetRolesOfUser")

	userRoleService.UserRoleRepository.GetRolesOfUser()
}

func (userRoleService *UserRoleService) AssignRoleToUser() {
	userRoleService.logger.Info("userRoleService -> AssignRoleToUser")

	userRoleService.UserRoleRepository.AssignRoleToUser()
}

func (userRoleService *UserRoleService) RemoveRoleFromUser() {
	userRoleService.logger.Info("userRoleService -> RemoveRoleFromUser")

	userRoleService.UserRoleRepository.RemoveRoleFromUser()
}

func (userRoleService *UserRoleService) CheckUserHasSingleRole() {
	userRoleService.logger.Info("userRoleService -> CheckUserHasSingleRole")

	userRoleService.UserRoleRepository.CheckUserHasSingleRole()
}

func (userRoleService *UserRoleService) CheckUserHasAnyRoles() {
	userRoleService.logger.Info("userRoleService -> CheckUserHasAnyRoles")

	userRoleService.UserRoleRepository.CheckUserHasAnyRoles()
}

func (userRoleService *UserRoleService) CheckUserHasAllRoles() {
	userRoleService.logger.Info("userRoleService -> CheckUserHasAllRoles")

	userRoleService.UserRoleRepository.CheckUserHasAllRoles()
}

func NewUserRoleService(UserRoleRepository repositories.UserRoleRepositoryInterface, logger *zap.Logger, serverConfig *config.ServerConfig) UserRoleServiceInterface {

	userRoleService := &UserRoleService{
		UserRoleRepository: UserRoleRepository,
		logger:             logger,
		serverConfig:       serverConfig,
	}

	return userRoleService

}
