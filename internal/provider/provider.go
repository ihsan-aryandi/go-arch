package provider

import (
	"go-arch/internal/app/handler"
	"go.uber.org/fx"
)

type Handlers struct {
	UserHandler *handler.UserHandler
}

func Modules(option fx.Option) fx.Option {
	return fx.Options(
		option,

		// Your dependencies go here
		userProviders(),

		// Register handlers
		fx.Provide(registerHandlers()),
	)
}

func registerHandlers() interface{} {
	return func(
		userHandler *handler.UserHandler,
	) *Handlers {

		// Register your handlers
		return &Handlers{
			UserHandler: userHandler,
		}
	}
}
