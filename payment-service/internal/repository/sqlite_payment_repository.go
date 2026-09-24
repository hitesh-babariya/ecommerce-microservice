package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"ecommerce-microservice/payment-service/internal/model"
)

type SQLitePaymentRepository struct {
	db *sql.DB
}

func NewSQLitePaymentRepository(db *sql.DB) *SQLitePaymentRepository {
	return &SQLitePaymentRepository{
		db: db,
	}
}

func (r *SQLitePaymentRepository) Create(
	ctx context.Context,
	payment model.Payment,
) (model.Payment, error) {

	query := `
		INSERT INTO payments (
			order_id,
			user_id,
			amount,
			status
		)
		VALUES (?, ?, ?, ?)
		RETURNING
			id,
			order_id,
			user_id,
			amount,
			status,
			created_at
	`

	var createdPayment model.Payment

	err := r.db.QueryRowContext(
		ctx,
		query,
		payment.OrderID,
		payment.UserID,
		payment.Amount,
		payment.Status,
	).Scan(
		&createdPayment.ID,
		&createdPayment.OrderID,
		&createdPayment.UserID,
		&createdPayment.Amount,
		&createdPayment.Status,
		&createdPayment.CreatedAt,
	)

	if err != nil {
		return model.Payment{}, fmt.Errorf("create payment: %w", err)
	}

	return createdPayment, nil
}

func (r *SQLitePaymentRepository) GetPaymentByID(
	ctx context.Context,
	id int64,
) (model.Payment, error) {

	query := `
		SELECT
			id,
			order_id,
			user_id,
			amount,
			status,
			created_at
		FROM payments
		WHERE id = ?
	`

	var payment model.Payment

	err := r.db.QueryRowContext(
		ctx,
		query,
		id,
	).Scan(
		&payment.ID,
		&payment.OrderID,
		&payment.UserID,
		&payment.Amount,
		&payment.Status,
		&payment.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.Payment{}, model.ErrPaymentNotFound
		}

		return model.Payment{}, fmt.Errorf("get payment: %w", err)
	}

	return payment, nil
}

func (r *SQLitePaymentRepository) UpdateStatus(
	ctx context.Context,
	id int64,
	status string,
) (model.Payment, error) {

	query := `
		UPDATE payments
		SET status = ?
		WHERE id = ?
	`

	result, err := r.db.ExecContext(
		ctx,
		query,
		status,
		id,
	)

	if err != nil {
		return model.Payment{}, fmt.Errorf("update payment status: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return model.Payment{}, fmt.Errorf(
			"check updated payment: %w",
			err,
		)
	}

	if rowsAffected == 0 {
		return model.Payment{}, model.ErrPaymentNotFound
	}

	return r.GetPaymentByID(ctx, id)
}
