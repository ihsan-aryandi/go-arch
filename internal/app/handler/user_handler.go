package handler

import (
	"go-arch/internal/app/service/userservice"
	"go-arch/internal/entity/request"
)

type UserHandler struct {
	userService *userservice.UserService
}

func NewUserHandler(userService *userservice.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) GetAll(ctx *Context) {
	users := h.userService.GetAllUsers()

	ctx.JSON(200, Map{
		"message": "Data retrieved successfully",
		"data":    users,
	})
	//apires.ResponseSuccessDataRetrieved(w, users, nil)
}

func (h *UserHandler) GetById(ctx *Context) {

}

func (h *UserHandler) Create(ctx *Context) {
	body := new(request.User)
	if err := ctx.ReadBody(body); err != nil {
		ctx.JSON(400, Map{
			"message": "Invalid request",
		})
	}

	h.userService.Create(body)

	ctx.JSON(200, Map{
		"message": "Data created successfully",
	})
}

func (h *UserHandler) UpdateById(ctx *Context) {

}

func (h *UserHandler) DeleteById(ctx *Context) {

}
