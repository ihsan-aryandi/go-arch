package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

/*
	Database interface
*/
type Database interface {
	/*
		OpenConnection is to make a new connection to database
	*/
	OpenConnection(driverName, dsn string) (*sql.DB, error)

	/*
		CloseConnection is to close a specific database connection
	*/
	CloseConnection(db *sql.DB)
}

type database struct {}

func NewDatabase() Database {
	return &database{}
}

func (*database) OpenConnection(driverName, dsn string) (*sql.DB, error) {
	db, err := sql.Open(driverName, dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (*database) CloseConnection(db *sql.DB) {
	err := db.Close()
	if err != nil {
		log.Fatal(err)
	}
}
