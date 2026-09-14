package model

import "time"

type OrderStatus string

const (
	OrderStatusPending   OrderStatus = "PENDING"
	OrderStatusPaid      OrderStatus = "PAID"
	OrderStatusCancelled OrderStatus = "CANCELLED"
)

type Order struct {
	ID        int64       `json:"id"`
	UserID    int64       `json:"userId"`
	ProductID int64       `json:"productId"`
	Quantity  int         `json:"quantity"`
	Price     float64     `json:"price"`
	Status    OrderStatus `json:"status"`
	CreatedAt time.Time   `json:"createdAt"`
}

type CreateOrderRequest struct {
	UserID    int64   `json:"userId" binding:"required"`
	ProductID int64   `json:"productId" binding:"required"`
	Quantity  int     `json:"quantity" binding:"required,min=1"`
	Price     float64 `json:"price" binding:"required,gt=0"`
}

type OrderResponse struct {
	ID        int64       `json:"id"`
	UserID    int64       `json:"userId"`
	ProductID int64       `json:"productId"`
	Quantity  int         `json:"quantity"`
	Price     float64     `json:"price"`
	Status    OrderStatus `json:"status"`
	CreatedAt time.Time   `json:"createdAt"`
}
