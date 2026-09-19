package model

import "errors"

var (
	ErrOrderNotFound      = errors.New("order not found")
	ErrInvalidOrderStatus = errors.New("invalid order status")
	ErrNoFieldsToUpdate   = errors.New("no fields provided for update")
)

type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
