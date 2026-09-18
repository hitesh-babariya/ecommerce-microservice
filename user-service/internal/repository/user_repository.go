package repository

import (
	"context"
	"ecommerce-microservice/user-service/internal/model"
)

type UserRepository interface {
	Create(ctx context.Context, user model.User) (model.User, error)
	GetUserByID(ctx context.Context, id int64) (model.User, error)
	UpdateUserByID(ctx context.Context, req model.UpdateUserRequest, id int64) (model.User, error)
	PatchUserByID(ctx context.Context, req model.PatchUserRequest, id int64) (model.User, error)
	DeleteUserByID(ctx context.Context, id int64) error
}
