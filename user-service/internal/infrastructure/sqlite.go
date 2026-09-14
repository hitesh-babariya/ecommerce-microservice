package infrastructure

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

func NewSQLite() (*sql.DB, error) {
	db, err := sql.Open("sqlite", "./ecommerce.db")
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
