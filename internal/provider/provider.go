package provider

import "go-arch/internal/app/handler"

type Modules struct {
	UserHandler *handler.UserHandler
}

func NewProvider(db string) *Modules {
	return &Modules{
		UserHandler: userHandlerProvider(db),
	}
}

//func NewProvider() *fx.App {
//	return fx.New(
//		/*
//			User
//		*/
//		fx.Provide(
//			"database",
//			repository.NewUserRepository,
//			userservice.NewUserService,
//			handler.NewUserHandler,
//		),
//	)
//}
