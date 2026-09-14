package repository

import (
	"ecommerce-microservice/user-service/internal/model"
)

type UserRepository interface {
	Create(user model.User) (model.User, error)
	GetUserByID(id int64) (model.User, error)
}
