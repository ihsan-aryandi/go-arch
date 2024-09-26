package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

type database struct{}

func NewDatabase() *database {
	return &database{}
}

/*
OpenConnection is to make a new connection to db
*/
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

/*
CloseConnection is to close a specific db connection
*/
func (*database) CloseConnection(db *sql.DB) {
	err := db.Close()
	if err != nil {
		log.Fatal(err)
	}
}
