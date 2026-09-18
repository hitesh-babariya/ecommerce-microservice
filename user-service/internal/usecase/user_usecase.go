package usecase

import (
	"context"
	"ecommerce-microservice/user-service/internal/model"
	"ecommerce-microservice/user-service/internal/repository"
)

type UserUsecase struct {
	repo repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) *UserUsecase {
	return &UserUsecase{
		repo: repo,
	}
}

func (u *UserUsecase) CreateUser(ctx context.Context, request model.CreateUserRequest) (model.User, error) {
	// Business logic

	user := model.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}
	return u.repo.Create(ctx, user)
}
func (u *UserUsecase) GetUserByID(ctx context.Context, id int64) (model.User, error) {
	// Business logic

	return u.repo.GetUserByID(ctx, id)
}
func (u *UserUsecase) UpdateUserByID(ctx context.Context, req model.UpdateUserRequest, id int64) (model.User, error) {
	// Business logic

	return u.repo.UpdateUserByID(ctx, req, id)
}
func (u *UserUsecase) PatchUserByID(ctx context.Context, req model.PatchUserRequest, id int64) (model.User, error) {
	// Business logic

	return u.repo.PatchUserByID(ctx, req, id)
}
func (u *UserUsecase) DeleteUserByID(ctx context.Context, id int64) error {
	// Business logic

	return u.repo.DeleteUserByID(ctx, id)
}
