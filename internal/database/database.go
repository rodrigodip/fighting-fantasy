package database

import (
	"api/internal/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql" //Driver
)

// Connection opens connection whit database
func Connection() (*sql.DB, error) {
	db, err := sql.Open("mysql", config.StrinConnect)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, err
}
