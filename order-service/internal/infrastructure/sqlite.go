package infrastructure

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

func NewSQLite() (*sql.DB, error) {

	db, err := sql.Open(
		"sqlite",
		"order-service/orders.db",
	)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS orders (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			user_id INTEGER NOT NULL,
			product_id INTEGER NOT NULL,
			quantity INTEGER NOT NULL,
			total_amount FLOAT NOT NULL,
			status TEXT NOT NULL,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		)
	`)

	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
