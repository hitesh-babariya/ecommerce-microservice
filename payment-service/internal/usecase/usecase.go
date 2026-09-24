package usecase

import (
	"context"
	"fmt"

	"ecommerce-microservice/payment-service/internal/model"
	"ecommerce-microservice/payment-service/internal/repository"
)

type PaymentUsecase struct {
	repo      repository.PaymentRepository
	processor PaymentProcessor
	publisher PaymentEventPublisher
}

func NewPaymentUsecase(
	repo repository.PaymentRepository,
	processor PaymentProcessor,
	publisher PaymentEventPublisher,
) *PaymentUsecase {

	return &PaymentUsecase{
		repo:      repo,
		processor: processor,
		publisher: publisher,
	}
}

func (u *PaymentUsecase) CreatePayment(
	ctx context.Context,
	req model.CreatePaymentRequest,
) (model.Payment, error) {

	// -------------------------
	// 1. Create PENDING payment
	// -------------------------

	payment := model.Payment{
		OrderID: req.OrderID,
		UserID:  req.UserID,
		Amount:  req.Amount,
		Status:  model.PaymentStatusPending,
	}

	createdPayment, err := u.repo.Create(
		ctx,
		payment,
	)

	if err != nil {
		return model.Payment{}, err
	}

	// -------------------------
	// 2. Process payment
	// -------------------------

	err = u.processor.ProcessPayment(
		ctx,
		createdPayment,
	)

	// -------------------------
	// 3. Payment FAILED
	// -------------------------

	if err != nil {

		failedPayment, updateErr := u.repo.UpdateStatus(
			ctx,
			createdPayment.ID,
			model.PaymentStatusFailed,
		)

		if updateErr != nil {
			return model.Payment{}, fmt.Errorf(
				"update failed payment status: %w",
				updateErr,
			)
		}

		if publishErr := u.publisher.PublishPaymentFailed(
			ctx,
			failedPayment,
		); publishErr != nil {
			return model.Payment{}, publishErr
		}

		return failedPayment, ErrPaymentProcessingFailed
	}

	// -------------------------
	// 4. Payment SUCCESS
	// -------------------------

	successPayment, err := u.repo.UpdateStatus(
		ctx,
		createdPayment.ID,
		model.PaymentStatusSuccess,
	)

	if err != nil {
		return model.Payment{}, fmt.Errorf(
			"update successful payment status: %w",
			err,
		)
	}

	// -------------------------
	// 5. Publish SUCCESS event
	// -------------------------

	if err := u.publisher.PublishPaymentSuccess(
		ctx,
		successPayment,
	); err != nil {
		return model.Payment{}, err
	}

	return successPayment, nil
}

func (u *PaymentUsecase) GetPaymentByID(
	ctx context.Context,
	id int64,
) (model.Payment, error) {

	return u.repo.GetPaymentByID(ctx, id)
}

func (u *PaymentUsecase) UpdatePaymentStatus(
	ctx context.Context,
	id int64,
	status string,
) (model.Payment, error) {

	switch status {

	case model.PaymentStatusPending,
		model.PaymentStatusSuccess,
		model.PaymentStatusFailed:

	default:
		return model.Payment{}, model.ErrInvalidPaymentStatus
	}

	return u.repo.UpdateStatus(
		ctx,
		id,
		status,
	)
}
