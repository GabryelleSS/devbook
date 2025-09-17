package db

import (
	"api/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // driver
)

func Connect() (*sql.DB, error) {
	db, erro := sql.Open("mysql", config.DBConnectionURL)

	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		db.Close()
		return nil, erro
	}

	return db, nil
}