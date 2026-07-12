package routers

import (
	"database/sql"

	"github.com/ayushWeb07/YT-Jwt-Auth-Golang/internal/config"
	"github.com/ayushWeb07/YT-Jwt-Auth-Golang/internal/controllers"
	"github.com/ayushWeb07/YT-Jwt-Auth-Golang/internal/repositories"
	"github.com/ayushWeb07/YT-Jwt-Auth-Golang/internal/services"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
)

type RouterInterface interface {
	Register(router *chi.Mux)
}

func RegisterRouters(db *sql.DB, logger *zap.Logger, serverConfig *config.ServerConfig) *chi.Mux {
	// create the router instance
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	// create all repositories
	userRepository := repositories.NewUserRepository(db, logger, serverConfig)
	roleRepository := repositories.NewRoleRepository(db, logger, serverConfig)
	userRoleRepository := repositories.NewUserRoleRepository(db, logger, serverConfig)

	// create all services
	userService := services.NewUserService(userRepository, logger, serverConfig)
	roleService := services.NewRoleService(roleRepository, logger, serverConfig)
	userRoleService := services.NewUserRoleService(userRoleRepository, logger, serverConfig)

	// create all controllers
	userController := controllers.NewUserController(userService, logger, serverConfig)
	roleController := controllers.NewRoleController(roleService, logger, serverConfig)
	userRoleController := controllers.NewUserRoleController(userRoleService, logger, serverConfig)

	// create all routers
	userRouter := NewUserRouter(userController, logger, serverConfig)
	roleRouter := NewRoleRouter(roleController, logger, serverConfig)
	userRoleRouter := NewUserRoleRouter(userRoleController, logger, serverConfig)

	// register all routers
	userRouter.Register(router)
	roleRouter.Register(router)
	userRoleRouter.Register(router)

	return router
}
