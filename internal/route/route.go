package route

import (
	"github.com/gorilla/mux"
	"go-arch/internal/core/router"
	"go-arch/internal/provider"
	"go-arch/internal/route/middleware"
)

func RegisterRoutes(port string) error {
	return router.SetupRouter(port, routes)
}

func routes(r *mux.Router) {
	providers := provider.NewProvider("d")

	r.HandleFunc("/greet", providers.UserHandler.GetUsers)

	r.Use(middleware.CORS)
}
