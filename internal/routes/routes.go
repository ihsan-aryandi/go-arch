package routes

import (
	"go-arch/internal/app/handler"
	"go-arch/internal/provider"
	"go-arch/internal/routes/middleware"
	"go-arch/pkg/router"
)

func Register(r *router.Engine, handlers *provider.Handlers) {
	// User endpoints
	userHandler := handlers.UserHandler

	r.Group("/user", router.GroupRouterMap{
		"GET":  middleware.Logger(handler.Fn(userHandler.GetAll)),
		"POST": middleware.Logger(handler.Fn(userHandler.Create)),
	})

	r.Group("/user/{id}", router.GroupRouterMap{
		"GET":    middleware.Logger(handler.Fn(userHandler.GetById)),
		"PUT":    handler.Fn(userHandler.Create),
		"DELETE": handler.Fn(userHandler.DeleteById),
	})

	// Middlewares
	r.Use(middleware.CORS)
}
