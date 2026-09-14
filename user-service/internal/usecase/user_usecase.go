package usecase

import (
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

func (u *UserUsecase) CreateUser(request model.CreateUserRequest) (model.User, error) {
	// Business logic

	user := model.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}
	return u.repo.Create(user)
}
func (u *UserUsecase) GetUserByID(id int64) (model.User, error) {
	// Business logic

	return u.repo.GetUserByID(id)
}
