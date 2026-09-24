package usecase

import (
	"context"

	"ecommerce-microservice/payment-service/internal/model"
)

type PaymentEventPublisher interface {
	PublishPaymentSuccess(
		ctx context.Context,
		payment model.Payment,
	) error

	PublishPaymentFailed(
		ctx context.Context,
		payment model.Payment,
	) error
}
