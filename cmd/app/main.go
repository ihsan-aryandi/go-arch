package main

import (
	"flag"
	"fmt"
	"go-arch/internal/config"
	"go-arch/internal/core/envloader"
	"go-arch/internal/route"
	"log"
)

func main() {
	var (
		cfg = &config.Conf{}
	)

	/*
		Parse Flags
	*/
	env := flag.String("env", "development", "environment (development, production)")
	flag.Parse()

	/*
		Load Config
	*/
	if err := envloader.LoadEnv(*env); err != nil {
		log.Fatalln(err.Error())
	}

	cfg = config.LoadConfig()

	/*
		Router
	*/
	log.Printf("Server started on port :%s\n", cfg.HTTP.Port)

	if err := route.RegisterRoutes(fmt.Sprintf(":%s", cfg.HTTP.Port)); err != nil {
		log.Fatalln(err.Error())
	}
}
