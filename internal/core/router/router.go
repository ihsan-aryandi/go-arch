package router

import (
	"github.com/gorilla/mux"
	"go-arch/internal/route"
	"net/http"
)

func SetupRouter(port string) error {
	r := mux.NewRouter()

	route.RegisterRoutes(r)

	return http.ListenAndServe(port, r)
}
