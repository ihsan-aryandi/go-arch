package provider

import (
	"go-arch/internal/app/handler"
	"go-arch/internal/app/repository"
	"go-arch/internal/app/service/userservice"
	"go.uber.org/fx"
)

func userProviders() fx.Option {
	return fx.Options(
		fx.Provide(
			repository.NewUserRepository,
			userservice.NewUserService,
			handler.NewUserHandler,
		),
	)
}
