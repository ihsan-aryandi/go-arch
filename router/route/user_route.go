package route

import (
	"github.com/gorilla/mux"
	"gofw/controller"
)

var (
	userController = controller.NewUserController()
)

func SetupUser(r *mux.Router) {
	r.HandleFunc("/user", userController.FindAllUsers).Methods("GET")
	r.HandleFunc("/user/{id}", userController.FindUserById).Methods("GET")
}