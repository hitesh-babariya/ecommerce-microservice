package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"ecommerce-microservice/user-service/internal/handler"
	framework "ecommerce-microservice/user-service/internal/infrastructure"
	"ecommerce-microservice/user-service/internal/repository"
	"ecommerce-microservice/user-service/internal/usecase"
)

func main() {

	// Framework / infrastructure
	db, err := framework.NewSQLite()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Repository
	userRepo := repository.NewSQLiteUserRepository(db)

	// Usecase
	userUsecase := usecase.NewUserUsecase(userRepo)

	// Handler
	userHandler := handler.NewUserHandler(userUsecase)

	// Gin
	router := gin.Default()

	router.POST("/users", userHandler.CreateUser)

	router.Run(":8080")
}
