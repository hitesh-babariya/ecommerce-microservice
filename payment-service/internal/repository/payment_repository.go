package repository

import (
	"context"

	"ecommerce-microservice/payment-service/internal/model"
)

type PaymentRepository interface {
	Create(
		ctx context.Context,
		payment model.Payment,
	) (model.Payment, error)

	GetPaymentByID(
		ctx context.Context,
		id int64,
	) (model.Payment, error)

	UpdateStatus(
		ctx context.Context,
		id int64,
		status string,
	) (model.Payment, error)
}
