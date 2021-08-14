package repo

import (
	"database/sql"
	"gofw/entity"
	"gofw/warehouse"
)

type exampleRepo struct {
	DB        *sql.DB
	TableName string
}

var users = []entity.User{
	{
		Id: 1,
		Username: "ihsan@gmail.com",
		Password: "ihsan123",
	},
	{
		Id: 2,
		Username: "jill@gmail.com",
		Password: "jill123",
	}}

func NewExampleRepo() *exampleRepo {
	return &exampleRepo{
		DB: warehouse.Conn.DB,
		TableName: "users"}
}

func (ur *exampleRepo) FindAll() []entity.User {
	return users
}

func (ur *exampleRepo) FindById(id int64) (user entity.User, isFound bool) {
	for _, item := range users {
		if item.Id == id {
			user = item
			isFound = true
			break
		}
	}
	return
}