package routers

import "github.com/go-chi/chi/v5"

type RouterInterface interface {
	Register(router *chi.Mux)
}
