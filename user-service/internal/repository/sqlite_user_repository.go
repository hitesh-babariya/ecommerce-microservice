package repository

import (
	"context"
	"database/sql"
	"ecommerce-microservice/user-service/internal/model"
)

type SQLiteUserRepository struct {
	db *sql.DB
}

func NewSQLiteUserRepository(db *sql.DB) *SQLiteUserRepository {
	return &SQLiteUserRepository{
		db: db,
	}
}

func (r *SQLiteUserRepository) Create(user model.User) (model.User, error) {

	ctx := context.Background()

	err := r.db.QueryRowContext(ctx, `
		INSERT INTO users (name, email, password)
		VALUES (?, ?, ?)
		RETURNING id, name, email, created_at
	`, user.Name, user.Email, user.Password,
	).Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt)

	return user, err
}

func (r *SQLiteUserRepository) GetUserByID(id int64) (user model.User, err error) {

	// Note In return variable user model.User, So it is named return variable so no need to declared again like below.
	//var user model.User
	// and err = r.db.QueryRowContext , should not used like err := r.db.QueryContext, because err variable already declared in (user model.User, err error)
	// so no need to declared again := means declared and assigned otherwise error will come : no new variables on left side of
	ctx := context.Background()

	err = r.db.QueryRowContext(ctx, `
		SELECT id, name, email, created_at 
		FROM users 
		WHERE id = ?
		`, id,
	).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.CreatedAt,
	)

	if err != nil {
		return model.User{}, err
	}

	return user, nil
}
