package routers

import (
	"github.com/ayushWeb07/YT-Jwt-Auth-Golang/internal/config"
	"github.com/ayushWeb07/YT-Jwt-Auth-Golang/internal/controllers"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

type UserRouter struct {
	UserController controllers.UserControllerInterface
	logger         *zap.Logger
	serverConfig   *config.ServerConfig
}

func (userRouter *UserRouter) Register(router *chi.Mux) {
	router.Route("/api/v1/users", func(r chi.Router) {
		r.Get("/", userRouter.UserController.GetAllUsers)
		r.Get("/{id}", userRouter.UserController.GetUserById)
		r.Post("/", userRouter.UserController.CreateUser)
		r.Put("/{id}", userRouter.UserController.UpdateUserById)
		r.Delete("/{id}", userRouter.UserController.DeleteUserById)
	})
}

func NewUserRouter(UserController controllers.UserControllerInterface, logger *zap.Logger, serverConfig *config.ServerConfig) RouterInterface {

	userRouter := &UserRouter{
		UserController: UserController,
		logger:         logger,
		serverConfig:   serverConfig,
	}

	return userRouter

}
