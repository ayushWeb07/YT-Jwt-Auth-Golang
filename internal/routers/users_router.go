package routers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ayushWeb07/YT-Jwt-Auth-Golang/internal/config"
	"github.com/ayushWeb07/YT-Jwt-Auth-Golang/internal/controllers"
	"github.com/ayushWeb07/YT-Jwt-Auth-Golang/internal/dtos"
	"github.com/ayushWeb07/YT-Jwt-Auth-Golang/internal/middlewares"
	"github.com/ayushWeb07/YT-Jwt-Auth-Golang/internal/utils"
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

		r.With(middlewares.DecodeAndValidateRequestParams[dtos.GetUserByIdParams](
			func(req *http.Request) (*dtos.GetUserByIdParams, *utils.AppError) {

				userId, err := strconv.Atoi(chi.URLParam(req, "id"))

				if err != nil {
					return nil, utils.BadRequestError(fmt.Sprintf("User id must be provided in integer: %s", err.Error()))
				}

				return &dtos.GetUserByIdParams{
					ID: userId,
				}, nil
			},
		)).Get("/{id}", userRouter.UserController.GetUserById)

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
