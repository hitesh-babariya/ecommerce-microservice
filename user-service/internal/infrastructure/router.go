package infrastructure

import (
	"github.com/gin-gonic/gin"

	"ecommerce-microservice/user-service/internal/handler"
)

func NewRouter(userHandler *handler.UserHandler) *gin.Engine {
	router := gin.Default()

	users := router.Group("/api/users")
	{
		users.POST("", userHandler.CreateUser)
		users.GET("/:id", userHandler.GetUserByID)
	}

	return router
}
