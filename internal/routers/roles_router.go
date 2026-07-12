package routers

import (
	"github.com/ayushWeb07/YT-Jwt-Auth-Golang/internal/config"
	"github.com/ayushWeb07/YT-Jwt-Auth-Golang/internal/controllers"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

type RoleRouter struct {
	RoleController controllers.RoleControllerInterface
	logger         *zap.Logger
	serverConfig   *config.ServerConfig
}

func (roleRouter *RoleRouter) Register(router *chi.Mux) {
	router.Route("/api/v1/roles", func(r chi.Router) {
		r.Get("/", roleRouter.RoleController.GetAllRoles)
		r.Get("/{id}", roleRouter.RoleController.GetRoleById)
		r.Post("/", roleRouter.RoleController.CreateRole)
		r.Put("/{id}", roleRouter.RoleController.UpdateRoleById)
		r.Delete("/{id}", roleRouter.RoleController.DeleteRoleById)
	})
}

func NewRoleRouter(RoleController controllers.RoleControllerInterface, logger *zap.Logger, serverConfig *config.ServerConfig) RouterInterface {

	roleRouter := &RoleRouter{
		RoleController: RoleController,
		logger:         logger,
		serverConfig:   serverConfig,
	}

	return roleRouter

}
