package data

import "database/sql"

type DataLayer struct {
	DataDB *sql.DB
}
