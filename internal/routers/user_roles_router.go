package routers

import (
	"github.com/ayushWeb07/YT-Jwt-Auth-Golang/internal/config"
	"github.com/ayushWeb07/YT-Jwt-Auth-Golang/internal/controllers"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

type UserRoleRouter struct {
	UserRoleController controllers.UserRoleControllerInterface
	logger             *zap.Logger
	serverConfig       *config.ServerConfig
}

func (userRoleRouter *UserRoleRouter) Register(router *chi.Mux) {
	router.Route("/api/v1/user-roles", func(r chi.Router) {
		r.Get("/user/{user_id}", userRoleRouter.UserRoleController.GetRolesOfUser)
		r.Post("/assign", userRoleRouter.UserRoleController.AssignRoleToUser)
		r.Post("/remove", userRoleRouter.UserRoleController.RemoveRoleFromUser)
		r.Get("/check-single", userRoleRouter.UserRoleController.CheckUserHasSingleRole)
		r.Get("/check-all", userRoleRouter.UserRoleController.CheckUserHasAllRoles)
		r.Get("/check-any", userRoleRouter.UserRoleController.CheckUserHasAnyRoles)
	})
}

func NewUserRoleRouter(UserRoleController controllers.UserRoleControllerInterface, logger *zap.Logger, serverConfig *config.ServerConfig) RouterInterface {

	userRoleRouter := &UserRoleRouter{
		UserRoleController: UserRoleController,
		logger:             logger,
		serverConfig:       serverConfig,
	}

	return userRoleRouter

}
