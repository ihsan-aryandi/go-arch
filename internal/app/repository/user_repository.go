package repository

import "database/sql"

type User struct {
	Id    sql.NullInt64
	Name  sql.NullString
	Email sql.NullString
}

type UserRepository struct {
	db string
}

var users = []*User{
	{
		Id:    sql.NullInt64{Int64: 1},
		Name:  sql.NullString{String: "John"},
		Email: sql.NullString{String: "john@gmail.com"},
	},
	{
		Id:    sql.NullInt64{Int64: 2},
		Name:  sql.NullString{String: "Jack"},
		Email: sql.NullString{String: "jack@gmail.com"},
	},
}

var sequence = int64(len(users) + 1)

func NewUserRepository(db string) *UserRepository {
	return &UserRepository{
		db,
	}
}

func (r *UserRepository) GetAll() []*User {
	return users
}

func (r *UserRepository) GetById(id int64) *User {
	for _, user := range users {
		if user.Id.Int64 == id {
			return user
		}
	}

	return nil
}

func (r *UserRepository) Create(user *User) (id int64) {
	user.Id.Int64 = sequence
	users = append(users, user)
	sequence++

	return user.Id.Int64
}

func (r *UserRepository) UpdateById(user *User) {
	for _, data := range users {
		if data.Id.Int64 == user.Id.Int64 {
			data = user
			return
		}
	}
}

func (r *UserRepository) DeleteById(user *User) {
	targetIndex := -1

	for i, data := range users {
		if data.Id.Int64 == user.Id.Int64 {
			targetIndex = i
			break
		}
	}

	if targetIndex == -1 {
		return
	}

	users = append(users[:targetIndex], users[targetIndex:len(users)-1]...)
}
