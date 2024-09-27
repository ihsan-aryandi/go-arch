package router

import (
	"github.com/gorilla/mux"
	"net/http"
)

type RouteHandler func(r *mux.Router)

func SetupRouter(port string, routes RouteHandler) error {
	r := mux.NewRouter()

	routes(r)

	return http.ListenAndServe(port, r)
}
