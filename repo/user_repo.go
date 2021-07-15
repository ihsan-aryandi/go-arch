package repo

import (
	"database/sql"
	"gofw/warehouse"
)

type UserRepo interface {
	FindAll()
}

type userRepo struct {
	DB        *sql.DB
	TableName string
}

func NewUserRepo() UserRepo {
	return &userRepo{
		DB: warehouse.DB,
		TableName: "users"}
}

func (ur *userRepo) FindAll() {
	// Some business here...
}