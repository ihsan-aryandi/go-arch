package service

import (
	"gofw/entity"
	"gofw/repo"
)

type exampleService struct {}

var (
	exampleRepo = repo.NewExampleRepo()
)

func NewExampleService() *exampleService {
	return &exampleService{}
}

func (*exampleService) FindAllUsers() []entity.User {
	return exampleRepo.FindAll()
}

func (*exampleService) FindUserById(id int64) (entity.User, bool) {
	return exampleRepo.FindById(id)
}