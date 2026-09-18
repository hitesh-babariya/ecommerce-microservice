package model

import "errors"

var (
	ErrUserNotFound       = errors.New("user not found")
	ErrEmailAlreadyExist  = errors.New("email already exists")
	ErrInvalidCredentials = errors.New("invalid credentials")
)
