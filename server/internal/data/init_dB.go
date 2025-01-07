package data

import (
	"database/sql"
	"os"
)

func openConn() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		return db, err
	}
	return db, nil
}

func initTables(db *sql.DB) error {
	data, err := os.ReadFile("./internal/data/tables.sql")
	if err != nil {
		return err
	}
	_, err = db.Exec(string(data))
	if err != nil {
		return err
	}
	return nil
}

func InitDb() (*sql.DB, error) {
	db, err := openConn()
	if err != nil {
		return db, err
	}
	err = initTables(db)
	if err != nil {
		return db, err
	}
	return db, nil
}
