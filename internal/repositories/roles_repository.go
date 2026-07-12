package repositories

import (
	"database/sql"

	"github.com/ayushWeb07/YT-Jwt-Auth-Golang/internal/config"
	"go.uber.org/zap"
)

type RoleRepositoryInterface interface {
	CreateRole()
	GetAllRoles()
	GetRoleById()
	UpdateRoleById()
	DeleteRoleById()
}

type RoleRepository struct {
	db           *sql.DB
	logger       *zap.Logger
	serverConfig *config.ServerConfig
}

func (roleRepository *RoleRepository) CreateRole() {
	roleRepository.logger.Info("roleRepository -> CreateRole")
}

func (roleRepository *RoleRepository) GetAllRoles() {
	roleRepository.logger.Info("roleRepository -> GetAllRoles")
}

func (roleRepository *RoleRepository) GetRoleById() {
	roleRepository.logger.Info("roleRepository -> GetRoleById")
}

func (roleRepository *RoleRepository) UpdateRoleById() {
	roleRepository.logger.Info("roleRepository -> UpdateRoleById")
}

func (roleRepository *RoleRepository) DeleteRoleById() {
	roleRepository.logger.Info("roleRepository -> DeleteRoleById")
}

func NewRoleRepository(db *sql.DB, logger *zap.Logger, serverConfig *config.ServerConfig) RoleRepositoryInterface {

	roleRepository := &RoleRepository{
		db:           db,
		logger:       logger,
		serverConfig: serverConfig,
	}

	return roleRepository

}
