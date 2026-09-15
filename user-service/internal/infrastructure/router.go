package infrastructure

import (
	"github.com/gin-gonic/gin"

	"ecommerce-microservice/user-service/internal/handler"
)

func NewRouter(userHandler *handler.UserHandler) *gin.Engine {
	router := gin.Default()

	// For Old client Versioning
	v1 := router.Group("/api/v1/users")
	{
		v1.POST("", userHandler.CreateUserV1)
		v1.GET("/:id", userHandler.GetUserByIDV1)
	}

	// For New client Versioning
	v2 := router.Group("/api/v2/users")
	{
		v2.POST("", userHandler.CreateUserV2)
		v2.GET("/:id", userHandler.GetUserByIDV2)
	}

	return router
}
