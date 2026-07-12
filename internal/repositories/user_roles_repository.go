package repositories

import (
	"database/sql"

	"github.com/ayushWeb07/YT-Jwt-Auth-Golang/internal/config"
	"go.uber.org/zap"
)

type UserRoleRepositoryInterface interface {
	GetRolesOfUser()
	AssignRoleToUser()
	RemoveRoleFromUser()
	CheckUserHasSingleRole()
	CheckUserHasAnyRoles()
	CheckUserHasAllRoles()
}

type UserRoleRepository struct {
	db           *sql.DB
	logger       *zap.Logger
	serverConfig *config.ServerConfig
}

func (userRoleRepository *UserRoleRepository) GetRolesOfUser() {
	userRoleRepository.logger.Info("userRoleRepository -> GetRolesOfUser")
}

func (userRoleRepository *UserRoleRepository) AssignRoleToUser() {
	userRoleRepository.logger.Info("userRoleRepository -> AssignRoleToUser")
}

func (userRoleRepository *UserRoleRepository) RemoveRoleFromUser() {
	userRoleRepository.logger.Info("userRoleRepository -> RemoveRoleFromUser")
}

func (userRoleRepository *UserRoleRepository) CheckUserHasSingleRole() {
	userRoleRepository.logger.Info("userRoleRepository -> CheckUserHasSingleRole")
}

func (userRoleRepository *UserRoleRepository) CheckUserHasAnyRoles() {
	userRoleRepository.logger.Info("userRoleRepository -> CheckUserHasAnyRoles")
}

func (userRoleRepository *UserRoleRepository) CheckUserHasAllRoles() {
	userRoleRepository.logger.Info("userRoleRepository -> CheckUserHasAllRoles")
}

func NewUserRoleRepository(db *sql.DB, logger *zap.Logger, serverConfig *config.ServerConfig) UserRoleRepositoryInterface {

	userRoleRepository := &UserRoleRepository{
		db:           db,
		logger:       logger,
		serverConfig: serverConfig,
	}

	return userRoleRepository

}
