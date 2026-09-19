package usecase

import "context"

type UserServiceClient interface {
	GetUserByID(ctx context.Context, userID int64) (User, error)
}

type User struct {
	ID        int64
	Name      string
	Email     string
	CreatedAt string
}
