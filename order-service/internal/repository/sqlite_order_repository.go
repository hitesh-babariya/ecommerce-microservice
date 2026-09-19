package repository

import (
	"context"
	"database/sql"
	"errors"
	"strings"

	"ecommerce-microservice/order-service/internal/model"
)

type SQLiteOrderRepository struct {
	db *sql.DB
}

func NewSQLiteOrderRepository(db *sql.DB) *SQLiteOrderRepository {
	return &SQLiteOrderRepository{
		db: db,
	}
}

func (r *SQLiteOrderRepository) Create(
	ctx context.Context,
	order model.Order,
) (model.Order, error) {

	err := r.db.QueryRowContext(
		ctx,
		`
		INSERT INTO orders
			(user_id, product_id, quantity, total_amount, status)
		VALUES
			(?, ?, ?, ?, ?)
		RETURNING
			id,
			user_id,
			product_id,
			quantity,
			total_amount,
			status,
			created_at
		`,
		order.UserID,
		order.ProductID,
		order.Quantity,
		order.TotalAmount,
		order.Status,
	).Scan(
		&order.ID,
		&order.UserID,
		&order.ProductID,
		&order.Quantity,
		&order.TotalAmount,
		&order.Status,
		&order.CreatedAt,
	)

	if err != nil {
		return model.Order{}, err
	}

	return order, nil
}

func (r *SQLiteOrderRepository) GetOrderByID(
	ctx context.Context,
	id int64,
) (model.Order, error) {

	var order model.Order

	err := r.db.QueryRowContext(
		ctx,
		`
		SELECT
			id,
			user_id,
			product_id,
			quantity,
			total_amount,
			status,
			created_at
		FROM orders
		WHERE id = ?
		`,
		id,
	).Scan(
		&order.ID,
		&order.UserID,
		&order.ProductID,
		&order.Quantity,
		&order.TotalAmount,
		&order.Status,
		&order.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.Order{}, model.ErrOrderNotFound
		}

		return model.Order{}, err
	}

	return order, nil
}

func (r *SQLiteOrderRepository) UpdateOrderByID(
	ctx context.Context,
	req model.UpdateOrderRequest,
	id int64,
) (model.Order, error) {

	query := `
		UPDATE orders
		SET
			user_id = ?,
			product_id = ?,
			quantity = ?,
			total_amount = ?,
			status = ?
		WHERE id = ?
	`

	result, err := r.db.ExecContext(
		ctx,
		query,
		req.UserID,
		req.ProductID,
		req.Quantity,
		req.TotalAmount,
		req.Status,
		id,
	)

	if err != nil {
		return model.Order{}, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return model.Order{}, err
	}

	if rowsAffected == 0 {
		return model.Order{}, model.ErrOrderNotFound
	}

	return r.GetOrderByID(ctx, id)
}

func (r *SQLiteOrderRepository) PatchOrderByID(
	ctx context.Context,
	req model.PatchOrderRequest,
	id int64,
) (model.Order, error) {

	query := "UPDATE orders SET "

	args := []any{}
	updates := []string{}

	if req.UserID != nil {
		updates = append(updates, "user_id = ?")
		args = append(args, *req.UserID)
	}

	if req.ProductID != nil {
		updates = append(updates, "product_id = ?")
		args = append(args, *req.ProductID)
	}

	if req.Quantity != nil {
		updates = append(updates, "quantity = ?")
		args = append(args, *req.Quantity)
	}

	if req.TotalAmount != nil {
		updates = append(updates, "total_amount = ?")
		args = append(args, *req.TotalAmount)
	}

	if req.Status != nil {
		updates = append(updates, "status = ?")
		args = append(args, *req.Status)
	}

	if len(updates) == 0 {
		return model.Order{}, model.ErrNoFieldsToUpdate
	}

	query += strings.Join(updates, ", ")
	query += " WHERE id = ?"

	args = append(args, id)

	result, err := r.db.ExecContext(
		ctx,
		query,
		args...,
	)

	if err != nil {
		return model.Order{}, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return model.Order{}, err
	}

	if rowsAffected == 0 {
		return model.Order{}, model.ErrOrderNotFound
	}

	return r.GetOrderByID(ctx, id)
}

func (r *SQLiteOrderRepository) DeleteOrderByID(
	ctx context.Context,
	id int64,
) error {

	result, err := r.db.ExecContext(
		ctx,
		`DELETE FROM orders WHERE id = ?`,
		id,
	)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return model.ErrOrderNotFound
	}

	return nil
}
