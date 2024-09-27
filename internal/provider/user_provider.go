package provider

import (
	"go-arch/internal/app/handler"
	"go-arch/internal/app/repository"
	"go-arch/internal/app/service/userservice"
)

func userHandlerProvider(db string) *handler.UserHandler {
	userRepo := repository.NewUserRepository(db)
	userSvc := userservice.NewUserService(userRepo)

	return handler.NewUserHandler(userSvc)
}
