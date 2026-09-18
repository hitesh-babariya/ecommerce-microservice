package repository

import (
	"context"
	"database/sql"
	"ecommerce-microservice/user-service/internal/model"
	"errors"
	"fmt"
	"strings"
)

type SQLiteUserRepository struct {
	db *sql.DB
}

func NewSQLiteUserRepository(db *sql.DB) *SQLiteUserRepository {
	return &SQLiteUserRepository{
		db: db,
	}
}

func (r *SQLiteUserRepository) Create(ctx context.Context, user model.User) (model.User, error) {

	err := r.db.QueryRowContext(ctx, `
		INSERT INTO users (name, email, password)
		VALUES (?, ?, ?)
		RETURNING id, name, email, created_at
	`, user.Name, user.Email, user.Password,
	).Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt)

	return user, err
}

func (r *SQLiteUserRepository) GetUserByID(ctx context.Context, id int64) (user model.User, err error) {

	// Note In return variable user model.User, So it is named return variable so no need to declared again like below.
	//var user model.User
	// and err = r.db.QueryRowContext , should not used like err := r.db.QueryContext, because err variable already declared in (user model.User, err error)
	// so no need to declared again := means declared and assigned otherwise error will come : no new variables on left side of

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

func (r *SQLiteUserRepository) UpdateUserByID(ctx context.Context, req model.UpdateUserRequest, id int64) (model.User, error) {

	query := `UPDATE users SET name=?, email=?, password=? WHERE id=?`

	result, err := r.db.ExecContext(ctx, query, req.Name, req.Email, req.Password, id)
	if err != nil {
		return model.User{}, errors.New("Faild to update user")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return model.User{}, errors.New("Failed to get affected rows")
	}

	if rowsAffected == 0 {
		return model.User{}, errors.New("User not found")
	}

	return model.User{
		ID:       id,
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}, nil

}
func (r *SQLiteUserRepository) PatchUserByID(
	ctx context.Context,
	req model.PatchUserRequest,
	id int64,
) (model.User, error) {

	query := "UPDATE users SET "
	args := []any{}
	updates := []string{}

	if req.Name != nil {
		updates = append(updates, "name = ?")
		args = append(args, *req.Name)
	}

	if req.Email != nil {
		updates = append(updates, "email = ?")
		args = append(args, *req.Email)
	}

	if req.Password != nil {
		updates = append(updates, "password = ?")
		args = append(args, *req.Password)
	}

	if len(updates) == 0 {
		return model.User{}, errors.New("no fields provided for update")
	}

	query += strings.Join(updates, ", ")
	query += " WHERE id = ?"

	args = append(args, id)

	_, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return model.User{}, errors.New("failed to update user")
	}

	// Get complete updated user
	return r.GetUserByID(ctx, id)
}

func (r *SQLiteUserRepository) DeleteUserByID(
	ctx context.Context,
	id int64,
) error {

	query := `DELETE FROM users WHERE id = ?`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("delete user: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("check deleted user: %w", err)
	}

	if rowsAffected == 0 {
		return model.ErrUserNotFound
	}

	return nil
}
