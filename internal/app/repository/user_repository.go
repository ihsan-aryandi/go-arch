package repository

import "database/sql"

type User struct {
	Name  sql.NullString
	Email sql.NullString
}

type UserRepository struct {
	db string
}

func NewUserRepository(db string) *UserRepository {
	return &UserRepository{
		db,
	}
}
