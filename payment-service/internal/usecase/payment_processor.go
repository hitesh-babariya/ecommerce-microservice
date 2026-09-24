package usecase

import (
	"context"
	"errors"
	"time"

	"ecommerce-microservice/payment-service/internal/model"
)

var ErrPaymentProcessingFailed = errors.New(
	"payment processing failed",
)

type PaymentProcessor interface {
	ProcessPayment(
		ctx context.Context,
		payment model.Payment,
	) error
}

type MockPaymentProcessor struct{}

func NewMockPaymentProcessor() *MockPaymentProcessor {
	return &MockPaymentProcessor{}
}

func (p *MockPaymentProcessor) ProcessPayment(
	ctx context.Context,
	payment model.Payment,
) error {

	// Simulate communication with a payment gateway.
	select {
	case <-time.After(2 * time.Second):
		return nil

	case <-ctx.Done():
		return ctx.Err()
	}
}
