package infrastructure

import (
	"github.com/gin-gonic/gin"

	"ecommerce-microservice/payment-service/internal/handler"
)

func NewRouter(paymentHandler *handler.PaymentHandler) *gin.Engine {

	router := gin.Default()

	payments := router.Group("/api/v1/payments")
	{
		payments.POST("", paymentHandler.CreatePayment)

		payments.GET(
			"/:id",
			paymentHandler.GetPaymentByID,
		)

		payments.PATCH(
			"/:id/status",
			paymentHandler.UpdatePaymentStatus,
		)
	}

	return router
}
