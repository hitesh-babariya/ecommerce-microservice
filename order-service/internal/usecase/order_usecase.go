package usecase

import (
	"context"

	"ecommerce-microservice/order-service/internal/model"
	"ecommerce-microservice/order-service/internal/repository"
)

type OrderUsecase struct {
	repo       repository.OrderRepository
	userClient UserServiceClient
}

func NewOrderUsecase(
	repo repository.OrderRepository,
	userClient UserServiceClient,
) *OrderUsecase {

	return &OrderUsecase{
		repo:       repo,
		userClient: userClient,
	}
}

func (u *OrderUsecase) CreateOrder(
	ctx context.Context,
	req model.CreateOrderRequest,
) (model.Order, error) {

	// 1. Verify that user exists
	_, err := u.userClient.GetUserByID(
		ctx,
		req.UserID,
	)

	if err != nil {
		return model.Order{}, err
	}

	// 2. User exists, create order
	order := model.Order{
		UserID:      req.UserID,
		ProductID:   req.ProductID,
		Quantity:    req.Quantity,
		TotalAmount: req.TotalAmount,
		Status:      model.OrderStatusPending,
	}

	return u.repo.Create(ctx, order)
}

func (u *OrderUsecase) GetOrderByID(
	ctx context.Context,
	id int64,
) (model.Order, error) {

	return u.repo.GetOrderByID(ctx, id)
}

func (u *OrderUsecase) UpdateOrderByID(
	ctx context.Context,
	req model.UpdateOrderRequest,
	id int64,
) (model.Order, error) {

	return u.repo.UpdateOrderByID(ctx, req, id)
}

func (u *OrderUsecase) PatchOrderByID(
	ctx context.Context,
	req model.PatchOrderRequest,
	id int64,
) (model.Order, error) {

	return u.repo.PatchOrderByID(ctx, req, id)
}

func (u *OrderUsecase) DeleteOrderByID(
	ctx context.Context,
	id int64,
) error {

	return u.repo.DeleteOrderByID(ctx, id)
}
func (u *OrderUsecase) UpdateOrderStatus(
	ctx context.Context,
	id int64,
	status string,
) (model.Order, error) {

	switch status {

	case model.OrderStatusPending,
		model.OrderStatusPaid,
		model.OrderStatusCancelled:

	default:
		return model.Order{}, model.ErrInvalidOrderStatus
	}

	return u.repo.UpdateStatus(
		ctx,
		id,
		status,
	)
}
