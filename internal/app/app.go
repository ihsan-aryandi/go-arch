package app

import (
	"flag"
	"fmt"
	"go-arch/internal/core/config"
)

func Start() {
	/*
		Config
	*/
	env := flag.String("env", "development", "environment (development, production)")
	flag.Parse()

	fmt.Println(*env)

	config.LoadConfig(*env)

	/*
		Router
	*/
	//if err := router.SetupRouter(""); err != nil {
	//	log.Fatal(err)
	//}
}
