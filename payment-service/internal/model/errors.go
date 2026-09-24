package model

import "errors"

var (
	ErrPaymentNotFound      = errors.New("payment not found")
	ErrInvalidPaymentStatus = errors.New("invalid payment status")
	ErrPaymentAlreadyExists = errors.New("payment already exists")
)

type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
