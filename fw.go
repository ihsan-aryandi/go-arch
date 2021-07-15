package main

import (
	"database/sql"
	"gofw/env"
	"gofw/internal/database"
	"gofw/router"
	"gofw/util"
	"gofw/warehouse"
	"log"
	"os"
)

var (
	db = database.NewDatabase()
	envReader = env.NewEnv()
	muxRouter = router.NewMuxRouter()
)

/*
	StartApplication is an application starter function
*/
func StartApplication() {
	var err error

	/*
		Read env
	*/
	err = envReader.ReadEnv()
	if err != nil {
		log.Fatal(err)
	}

	/*
		Connect to main DB
	*/
	warehouse.DB, err = connectDB()
	if err != nil {
		log.Fatal(err)
	}

	/*
		Starts Router
	*/
	log.Fatal(muxRouter.Listen(os.Getenv("PORT")))
}

func connectDB() (*sql.DB, error) {
	driverName := os.Getenv("DB_DRIVER_NAME")

	dsn := util.NewDSNCreator()
	dsn.DBName(os.Getenv("DB_NAME"))
	dsn.User(os.Getenv("DB_USER"))
	dsn.Password(os.Getenv("DB_PASSWORD"))
	dsn.Host(os.Getenv("DB_HOST"))
	dsn.SSLMode(os.Getenv("DB_SSL_MODE"))
	dsn.Schema(os.Getenv("DB_SCHEMA"))

	return db.OpenConnection(driverName, dsn.Create())
}