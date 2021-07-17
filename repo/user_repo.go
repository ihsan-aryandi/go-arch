package repo

import (
	"database/sql"
	"gofw/entity"
	"gofw/warehouse"
)

type UserRepo interface {
	FindAll() []entity.User
	FindById(id int64) (entity.User, bool)
}

type userRepo struct {
	DB        *sql.DB
	TableName string
}

var users []entity.User

func NewUserRepo() UserRepo {
	users = append(users, entity.User{
		Id: 1,
		Username: "ihsan@gmail.com",
		Password: "ihsan123",
	})
	users = append(users, entity.User{
		Id: 2,
		Username: "jill@gmail.com",
		Password: "jill123",
	})

	return &userRepo{
		DB: warehouse.DB,
		TableName: "users"}
}

func (ur *userRepo) FindAll() (users []entity.User) {
	return users
}

func (ur *userRepo) FindById(id int64) (user entity.User, isFound bool) {
	for _, item := range users {
		if item.Id == id {
			user = item
			isFound = true
			break
		}
	}
	return
}