package model

import "time"

const (
	OrderStatusPending   = "PENDING"
	OrderStatusPaid      = "PAID"
	OrderStatusCancelled = "CANCELLED"
)

type Order struct {
	ID          int64     `json:"id"`
	UserID      int64     `json:"userId"`
	ProductID   int64     `json:"productId"`
	Quantity    int       `json:"quantity"`
	TotalAmount float64   `json:"totalAmount"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
}

type CreateOrderRequest struct {
	UserID      int64   `json:"userId" binding:"required,gt=0"`
	ProductID   int64   `json:"productId" binding:"required,gt=0"`
	Quantity    int     `json:"quantity" binding:"required,gt=0"`
	TotalAmount float64 `json:"totalAmount" binding:"required,gt=0"`
}

type UpdateOrderRequest struct {
	UserID      int64   `json:"userId" binding:"required,gt=0"`
	ProductID   int64   `json:"productId" binding:"required,gt=0"`
	Quantity    int     `json:"quantity" binding:"required,gt=0"`
	TotalAmount float64 `json:"totalAmount" binding:"required,gt=0"`
	Status      string  `json:"status" binding:"required"`
}

type PatchOrderRequest struct {
	UserID      *int64   `json:"userId"`
	ProductID   *int64   `json:"productId"`
	Quantity    *int     `json:"quantity"`
	TotalAmount *float64 `json:"totalAmount"`
	Status      *string  `json:"status"`
}

type OrderResponse struct {
	ID          int64     `json:"id"`
	UserID      int64     `json:"userId"`
	ProductID   int64     `json:"productId"`
	Quantity    int       `json:"quantity"`
	TotalAmount float64   `json:"totalAmount"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
}

type DeleteOrderResponse struct {
	ID      int64  `json:"id"`
	Message string `json:"message"`
}
