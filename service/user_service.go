package service

import (
	"gofw/entity"
	"gofw/repo"
)

type UserService interface {
	FindAllUsers() []entity.User
	FindUserById(id int64) (entity.User, bool)
}

type userService struct {}

var (
	userRepo = repo.NewUserRepo()
)

func NewUserService() UserService {
	return &userService{}
}

func (*userService) FindAllUsers() []entity.User {
	return userRepo.FindAll()
}

func (*userService) FindUserById(id int64) (entity.User, bool) {
	return userRepo.FindById(id)
}