package route

import (
	"github.com/gorilla/mux"
	"net/http"
)

func RegisterUserRoutes(r *mux.Router) {
	r.HandleFunc("/user", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("User Page"))
	}).Methods("GET")
}