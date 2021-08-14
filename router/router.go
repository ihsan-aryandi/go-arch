package router

import (
	"github.com/gorilla/mux"
	"gofw/router/route"
)

func SetupRoutes(r *mux.Router) {
	/*
		Routes Setup
	*/
	route.SetupExampleRoute(r)

	/*
		Global Middleware Setup
	*/
	r.Use()
}