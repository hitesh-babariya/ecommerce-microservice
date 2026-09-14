package infrastructure

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

func NewSQLite() (*sql.DB, error) {
	db, err := sql.Open("sqlite", "user-service/users.db")
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	_, err = db.Exec(` CREATE TABLE IF NOT EXISTS users ( id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT NOT NULL, email TEXT NOT NULL UNIQUE, password TEXT NOT NULL, created_at DATETIME DEFAULT CURRENT_TIMESTAMP ) `)

	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
