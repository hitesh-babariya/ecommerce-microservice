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
		v1.PUT("/:id", userHandler.UpdateUserByIDV1)  // Update By PUT Whole User Data of a User
		v1.PATCH("/:id", userHandler.PatchUserByIDV1) // Update By PATCH Partially update data of a User
		v1.DELETE("/:id", userHandler.DeleteUserByIDV1)
	}

	// For New client Versioning
	v2 := router.Group("/api/v2/users")
	{
		v2.POST("", userHandler.CreateUserV2)
		v2.GET("/:id", userHandler.GetUserByIDV2)
		v2.PUT("/:id", userHandler.UpdateUserByIDV2)  // Update By PUT Whole User Data of a User
		v2.PATCH("/:id", userHandler.PatchUserByIDV2) // Update By PATCH Partially update data of a User
	}

	return router
}
