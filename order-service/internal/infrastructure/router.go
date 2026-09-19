package infrastructure

import (
	"github.com/gin-gonic/gin"

	"ecommerce-microservice/order-service/internal/handler"
)

func NewRouter(orderHandler *handler.OrderHandler) *gin.Engine {

	router := gin.Default()

	v1 := router.Group("/api/v1/orders")
	{
		v1.POST("", orderHandler.CreateOrderV1)
		v1.GET("/:id", orderHandler.GetOrderByIDV1)
		v1.PUT("/:id", orderHandler.UpdateOrderByIDV1)
		v1.PATCH("/:id", orderHandler.PatchOrderByIDV1)
		v1.DELETE("/:id", orderHandler.DeleteOrderByIDV1)
	}

	return router
}
