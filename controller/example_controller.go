package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type exampleController struct {}

func NewExampleController() *exampleController {
	return &exampleController{}
}

func (uc *exampleController) FindAllUsers(w http.ResponseWriter, r *http.Request) {
	users := exampleService.FindAllUsers()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	_ = json.NewEncoder(w).Encode(users)
}

func (uc *exampleController) FindUserById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	varId := mux.Vars(r)["id"]
	id, err := strconv.Atoi(varId)
	if err != nil {
		_= json.NewEncoder(w).Encode(map[string]string{
			"message": "id is invalid",
		})
		return
	}

	user, isFound := exampleService.FindUserById(int64(id))
	if !isFound {
		_ = json.NewEncoder(w).Encode(nil)
		return
	}

	_ = json.NewEncoder(w).Encode(user)
}