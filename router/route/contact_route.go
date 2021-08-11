package route

import (
	"github.com/gorilla/mux"
	"net/http"
)

func SetupContact(r *mux.Router) {
	r.HandleFunc("/contact", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Contact Page"))
	}).Methods("GET")
}