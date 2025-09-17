package db

import (
	"api/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // driver
)

func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", config.DBConnectionURL)

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}