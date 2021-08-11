package warehouse

import "database/sql"

type Connections struct {
	DB *sql.DB
}

var Conn = &Connections{}
