package main

import (
	"context"
	"flag"
	"fmt"
	"go-arch/internal/config"
	"go-arch/internal/provider"
	"go-arch/internal/routes"
	"go-arch/pkg/envloader"
	"go-arch/pkg/router"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"log"
)

func main() {
	var (
		cfg = &config.Conf{}
	)

	// Parse flags
	env := flag.String("env", "development", "environment (development, production)")
	flag.Parse()

	// Load env
	if err := envloader.Load(*env); err != nil {
		log.Fatalln(err.Error())
	}

	// Load config
	cfg = config.LoadConfig()

	// Init router
	r := router.NewRouter()

	// Init providers
	db := "jdbc://localhost:5000"
	supplies := fx.Supply(r, cfg, db)

	app := fx.New(
		provider.Modules(supplies),
		fx.WithLogger(func() fxevent.Logger {
			return fxevent.NopLogger
		}),
		fx.Invoke(routes.Register),
	)

	// Check if provider works
	if err := app.Start(context.Background()); err != nil {
		log.Fatalln("Dependency injection failed: ", err.Error())
	}

	// Start router
	log.Printf("Server started on port :%s\n", cfg.HTTP.Port)

	_ = r.Run(fmt.Sprintf(":%s", cfg.HTTP.Port))
}
