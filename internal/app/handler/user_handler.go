package handler

import (
	"go-arch/internal/app/service/userservice"
	"net/http"
)

type UserHandler struct {
	userService *userservice.UserService
}

func NewUserHandler(userService *userservice.UserService) *UserHandler {
	return &UserHandler{
		userService,
	}
}

func (h *UserHandler) GetUsers(res http.ResponseWriter, req *http.Request) {

}
