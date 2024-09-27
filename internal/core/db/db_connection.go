package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

type Instance struct{}

func NewDatabase() *Instance {
	return &Instance{}
}

// OpenConnection created new connection
func (*Instance) OpenConnection(driverName, dsn string) (*sql.DB, error) {
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

// CloseConnection closes current connection
func (*Instance) CloseConnection(db *sql.DB) {
	err := db.Close()
	if err != nil {
		log.Fatal(err)
	}
}
