package router

import (
	"github.com/gorilla/mux"
	"gofw/route"
)

func RegisterRoutes(r *mux.Router) {
	/*
		Routes Registration
	*/
	route.RegisterUserRoutes(r)
	route.RegisterContactRoutes(r)


	/*
		Global Middleware Registration
	*/
	r.Use()
}