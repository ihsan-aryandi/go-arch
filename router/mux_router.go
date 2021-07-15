package router

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type MuxRouters interface {
	Listen(port string) error
}

type muxRouter struct {
	MuxRouter *mux.Router
}

func NewMuxRouter() MuxRouters {
	r := mux.NewRouter()
	RegisterRoutes(r)

	return &muxRouter{MuxRouter: r}
}

func (r *muxRouter) Listen(port string) error {
	port = fmt.Sprintf(":%v", port)

	fmt.Println(fmt.Sprintf("Server started on port %v", port))

	return http.ListenAndServe(port, r.MuxRouter)
}
