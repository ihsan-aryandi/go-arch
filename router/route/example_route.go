package route

import (
	"github.com/gorilla/mux"
	"gofw/controller"
)

var (
	exampleController = controller.NewExampleController()
)

func SetupExampleRoute(r *mux.Router) {
	r.HandleFunc("/user", exampleController.FindAllUsers).Methods("GET")
	r.HandleFunc("/user/{id}", exampleController.FindUserById).Methods("GET")
}