package router

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Engine struct {
	router *mux.Router
}

type GroupRouterMap map[string]http.HandlerFunc

func NewRouter() *Engine {
	return &Engine{
		router: mux.NewRouter(),
	}
}

func (e *Engine) HandleFunc(path string, handlerFunc http.HandlerFunc) *mux.Router {
	e.router.HandleFunc(path, handlerFunc)
	return e.router
}

func (e *Engine) Use(middlewares ...mux.MiddlewareFunc) {
	e.router.Use(middlewares...)
}

func (e *Engine) Group(path string, routerMap GroupRouterMap) {
	methods := []string{http.MethodOptions}

	for method, _ := range routerMap {
		if method == http.MethodOptions {
			continue
		}

		methods = append(methods, method)
	}

	handlerFunc := func(writer http.ResponseWriter, request *http.Request) {
		fn := routerMap[request.Method]

		if fn != nil {
			fn(writer, request)
		}
	}

	e.router.HandleFunc(path, handlerFunc).Methods(methods...)
}

func (e *Engine) Router() *mux.Router {
	return e.router
}

func (e *Engine) Run(port string) error {
	return http.ListenAndServe(port, e.router)
}
