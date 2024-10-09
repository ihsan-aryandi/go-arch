package userservice

import (
	"database/sql"
	"go-arch/internal/app/repository"
	"go-arch/internal/entity/request"
	"go-arch/internal/entity/response"
)

type UserService struct {
	userRepository *repository.UserRepository
}

func NewUserService(userRepository *repository.UserRepository) *UserService {
	return &UserService{
		userRepository,
	}
}

func (s *UserService) GetAllUsers() []*response.User {
	users := s.userRepository.GetAll()

	var results []*response.User

	for _, user := range users {
		results = append(results, &response.User{
			ID:    user.Id.Int64,
			Name:  user.Name.String,
			Email: user.Email.String,
		})
	}

	return results
}

func (s *UserService) Create(user *request.User) {
	s.userRepository.Create(&repository.User{
		Name:  sql.NullString{String: user.Name},
		Email: sql.NullString{String: user.Email},
	})
}
