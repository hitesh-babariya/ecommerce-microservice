package repository

import (
	"context"

	"ecommerce-microservice/order-service/internal/model"
)

type OrderRepository interface {
	Create(
		ctx context.Context,
		order model.Order,
	) (model.Order, error)

	GetOrderByID(
		ctx context.Context,
		id int64,
	) (model.Order, error)

	UpdateOrderByID(
		ctx context.Context,
		req model.UpdateOrderRequest,
		id int64,
	) (model.Order, error)

	PatchOrderByID(
		ctx context.Context,
		req model.PatchOrderRequest,
		id int64,
	) (model.Order, error)

	DeleteOrderByID(
		ctx context.Context,
		id int64,
	) error
}
