package infrastructure

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

func NewSQLite() (*sql.DB, error) {

	db, err := sql.Open(
		"sqlite",
		"payment-service/payments.db",
	)
	if err != nil {
		return nil, fmt.Errorf("open sqlite database: %w", err)
	}

	if err := db.Ping(); err != nil {
		db.Close()

		return nil, fmt.Errorf("ping sqlite database: %w", err)
	}

	query := `
		CREATE TABLE IF NOT EXISTS payments (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			order_id INTEGER NOT NULL,
			user_id INTEGER NOT NULL,
			amount INTEGER NOT NULL,
			status TEXT NOT NULL,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		)
	`

	if _, err := db.Exec(query); err != nil {
		db.Close()

		return nil, fmt.Errorf("create payments table: %w", err)
	}

	return db, nil
}
