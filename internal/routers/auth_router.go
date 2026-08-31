package routers

import (
	"github.com/ayushWeb07/YT-Jwt-Auth-Golang/internal/config"
	"github.com/ayushWeb07/YT-Jwt-Auth-Golang/internal/controllers"
	"github.com/ayushWeb07/YT-Jwt-Auth-Golang/internal/dtos"
	"github.com/ayushWeb07/YT-Jwt-Auth-Golang/internal/middlewares"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

type AuthRouter struct {
	AuthController controllers.AuthControllerInterface
	logger         *zap.Logger
	serverConfig   *config.ServerConfig
}

func (authRouter *AuthRouter) Register(router *chi.Mux) {
	router.Route("/api/v1/auth", func(r chi.Router) {
		r.With(middlewares.DecodeAndValidateRequestBody[dtos.CreateUserPayload]).Post("/register", authRouter.AuthController.RegisterUser)
	})
}

func NewAuthRouter(authController controllers.AuthControllerInterface, logger *zap.Logger, serverConfig *config.ServerConfig) RouterInterface {

	authRouter := &AuthRouter{
		AuthController: authController,
		logger:         logger,
		serverConfig:   serverConfig,
	}

	return authRouter

}
