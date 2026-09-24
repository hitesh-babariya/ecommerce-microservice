package handler

import (
	"context"
	"errors"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"ecommerce-microservice/payment-service/internal/model"
	"ecommerce-microservice/payment-service/internal/usecase"
)

type PaymentHandler struct {
	paymentUsecase *usecase.PaymentUsecase
}

func NewPaymentHandler(
	paymentUsecase *usecase.PaymentUsecase,
) *PaymentHandler {

	return &PaymentHandler{
		paymentUsecase: paymentUsecase,
	}
}

func (h *PaymentHandler) CreatePayment(c *gin.Context) {

	var request model.CreatePaymentRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Code:    "INVALID_REQUEST",
			Message: "Invalid payment request",
		})
		return
	}

	ctx, cancel := context.WithTimeout(
		c.Request.Context(),
		2*time.Second,
	)
	defer cancel()

	payment, err := h.paymentUsecase.CreatePayment(
		ctx,
		request,
	)

	if err != nil {
		if errors.Is(
			err,
			usecase.ErrPaymentProcessingFailed,
		) {
			c.JSON(
				http.StatusPaymentRequired,
				model.ErrorResponse{
					Code:    "PAYMENT_FAILED",
					Message: "Payment failed",
				},
			)
			return
		}

		log.Println("create payment error:", err)

		c.JSON(
			http.StatusInternalServerError,
			model.ErrorResponse{
				Code:    "INTERNAL_SERVER_ERROR",
				Message: "Internal server error",
			},
		)

		return
	}

	c.JSON(http.StatusCreated, payment)
}

func (h *PaymentHandler) GetPaymentByID(c *gin.Context) {

	id, err := strconv.ParseInt(
		c.Param("id"),
		10,
		64,
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Code:    "INVALID_PAYMENT_ID",
			Message: "Invalid payment ID",
		})
		return
	}

	ctx, cancel := context.WithTimeout(
		c.Request.Context(),
		2*time.Second,
	)
	defer cancel()

	payment, err := h.paymentUsecase.GetPaymentByID(
		ctx,
		id,
	)

	if err != nil {

		if errors.Is(err, model.ErrPaymentNotFound) {
			c.JSON(http.StatusNotFound, model.ErrorResponse{
				Code:    "PAYMENT_NOT_FOUND",
				Message: "Payment not found",
			})
			return
		}

		log.Println("get payment error:", err)

		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Code:    "INTERNAL_SERVER_ERROR",
			Message: "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, payment)
}

func (h *PaymentHandler) UpdatePaymentStatus(c *gin.Context) {

	id, err := strconv.ParseInt(
		c.Param("id"),
		10,
		64,
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Code:    "INVALID_PAYMENT_ID",
			Message: "Invalid payment ID",
		})
		return
	}

	var request model.UpdatePaymentStatusRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Code:    "INVALID_REQUEST",
			Message: "Invalid payment status request",
		})
		return
	}

	ctx, cancel := context.WithTimeout(
		c.Request.Context(),
		2*time.Second,
	)
	defer cancel()

	payment, err := h.paymentUsecase.UpdatePaymentStatus(
		ctx,
		id,
		request.Status,
	)

	if err != nil {

		if errors.Is(err, model.ErrPaymentNotFound) {
			c.JSON(http.StatusNotFound, model.ErrorResponse{
				Code:    "PAYMENT_NOT_FOUND",
				Message: "Payment not found",
			})
			return
		}

		if errors.Is(err, model.ErrInvalidPaymentStatus) {
			c.JSON(http.StatusBadRequest, model.ErrorResponse{
				Code:    "INVALID_PAYMENT_STATUS",
				Message: "Invalid payment status",
			})
			return
		}

		log.Println("update payment status error:", err)

		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Code:    "INTERNAL_SERVER_ERROR",
			Message: "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, payment)
}
