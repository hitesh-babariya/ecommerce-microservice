package model

import "time"

const (
	PaymentStatusPending = "PENDING"
	PaymentStatusSuccess = "SUCCESS"
	PaymentStatusFailed  = "FAILED"
)

type Payment struct {
	ID        int64     `json:"id"`
	OrderID   int64     `json:"orderId"`
	UserID    int64     `json:"userId"`
	Amount    float64   `json:"amount"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
}

type CreatePaymentRequest struct {
	OrderID int64   `json:"orderId" binding:"required"`
	UserID  int64   `json:"userId" binding:"required"`
	Amount  float64 `json:"amount" binding:"required,gt=0"`
}

type UpdatePaymentStatusRequest struct {
	Status string `json:"status" binding:"required"`
}

type PaymentResponse struct {
	ID        int64     `json:"id"`
	OrderID   int64     `json:"orderId"`
	UserID    int64     `json:"userId"`
	Amount    float64   `json:"amount"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
}
