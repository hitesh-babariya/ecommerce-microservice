package handler

import (
	"context"
	"errors"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"ecommerce-microservice/order-service/internal/model"
	"ecommerce-microservice/order-service/internal/usecase"
)

type OrderHandler struct {
	orderUsecase *usecase.OrderUsecase
}

func NewOrderHandler(
	orderUsecase *usecase.OrderUsecase,
) *OrderHandler {

	return &OrderHandler{
		orderUsecase: orderUsecase,
	}
}

// POST /api/v1/orders
func (h *OrderHandler) CreateOrderV1(c *gin.Context) {

	var req model.CreateOrderRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Code:    "VALIDATION_ERROR",
			Message: "Invalid request data",
		})

		return
	}

	ctx, cancel := context.WithTimeout(
		c.Request.Context(),
		2*time.Second,
	)
	defer cancel()

	order, err := h.orderUsecase.CreateOrder(ctx, req)

	if err != nil {

		if errors.Is(err, model.ErrUserNotFound) {

			c.JSON(http.StatusNotFound, model.ErrorResponse{
				Code:    "USER_NOT_FOUND",
				Message: "User not found",
			})

			return
		}

		log.Println("Error:", err)

		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Code:    "INTERNAL_SERVER_ERROR",
			Message: "Internal server error",
		})

		return
	}

	response := model.OrderResponse{
		ID:          order.ID,
		UserID:      order.UserID,
		ProductID:   order.ProductID,
		Quantity:    order.Quantity,
		TotalAmount: order.TotalAmount,
		Status:      order.Status,
		CreatedAt:   order.CreatedAt,
	}

	c.JSON(http.StatusCreated, response)
}

// GET /api/v1/orders/:id
func (h *OrderHandler) GetOrderByIDV1(c *gin.Context) {

	id, err := strconv.ParseInt(
		c.Param("id"),
		10,
		64,
	)

	if err != nil {

		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Code:    "INVALID_ORDER_ID",
			Message: "Invalid order id",
		})

		return
	}

	ctx, cancel := context.WithTimeout(
		c.Request.Context(),
		2*time.Second,
	)
	defer cancel()

	order, err := h.orderUsecase.GetOrderByID(ctx, id)

	if err != nil {

		if errors.Is(err, model.ErrOrderNotFound) {

			c.JSON(http.StatusNotFound, model.ErrorResponse{
				Code:    "ORDER_NOT_FOUND",
				Message: "Order not found",
			})

			return
		}

		log.Println("Error:", err)

		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Code:    "INTERNAL_SERVER_ERROR",
			Message: "Internal server error",
		})

		return
	}

	c.JSON(http.StatusOK, model.OrderResponse{
		ID:          order.ID,
		UserID:      order.UserID,
		ProductID:   order.ProductID,
		Quantity:    order.Quantity,
		TotalAmount: order.TotalAmount,
		Status:      order.Status,
		CreatedAt:   order.CreatedAt,
	})
}

// PUT /api/v1/orders/:id
func (h *OrderHandler) UpdateOrderByIDV1(c *gin.Context) {

	id, err := strconv.ParseInt(
		c.Param("id"),
		10,
		64,
	)

	if err != nil {

		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Code:    "INVALID_ORDER_ID",
			Message: "Invalid order id",
		})

		return
	}

	var req model.UpdateOrderRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Code:    "VALIDATION_ERROR",
			Message: "Invalid request data",
		})

		return
	}

	ctx, cancel := context.WithTimeout(
		c.Request.Context(),
		2*time.Second,
	)
	defer cancel()

	order, err := h.orderUsecase.UpdateOrderByID(
		ctx,
		req,
		id,
	)

	if err != nil {

		if errors.Is(err, model.ErrOrderNotFound) {

			c.JSON(http.StatusNotFound, model.ErrorResponse{
				Code:    "ORDER_NOT_FOUND",
				Message: "Order not found",
			})

			return
		}

		log.Println("Error:", err)

		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Code:    "INTERNAL_SERVER_ERROR",
			Message: "Internal server error",
		})

		return
	}

	c.JSON(http.StatusOK, model.OrderResponse{
		ID:          order.ID,
		UserID:      order.UserID,
		ProductID:   order.ProductID,
		Quantity:    order.Quantity,
		TotalAmount: order.TotalAmount,
		Status:      order.Status,
		CreatedAt:   order.CreatedAt,
	})
}

// PATCH /api/v1/orders/:id
func (h *OrderHandler) PatchOrderByIDV1(c *gin.Context) {

	id, err := strconv.ParseInt(
		c.Param("id"),
		10,
		64,
	)

	if err != nil {

		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Code:    "INVALID_ORDER_ID",
			Message: "Invalid order id",
		})

		return
	}

	var req model.PatchOrderRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Code:    "VALIDATION_ERROR",
			Message: "Invalid request data",
		})

		return
	}

	ctx, cancel := context.WithTimeout(
		c.Request.Context(),
		2*time.Second,
	)
	defer cancel()

	order, err := h.orderUsecase.PatchOrderByID(
		ctx,
		req,
		id,
	)

	if err != nil {

		if errors.Is(err, model.ErrOrderNotFound) {

			c.JSON(http.StatusNotFound, model.ErrorResponse{
				Code:    "ORDER_NOT_FOUND",
				Message: "Order not found",
			})

			return
		}

		if errors.Is(err, model.ErrNoFieldsToUpdate) {

			c.JSON(http.StatusBadRequest, model.ErrorResponse{
				Code:    "NO_FIELDS_TO_UPDATE",
				Message: "No fields provided for update",
			})

			return
		}

		log.Println("Error:", err)

		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Code:    "INTERNAL_SERVER_ERROR",
			Message: "Internal server error",
		})

		return
	}

	c.JSON(http.StatusOK, model.OrderResponse{
		ID:          order.ID,
		UserID:      order.UserID,
		ProductID:   order.ProductID,
		Quantity:    order.Quantity,
		TotalAmount: order.TotalAmount,
		Status:      order.Status,
		CreatedAt:   order.CreatedAt,
	})
}

// DELETE /api/v1/orders/:id
func (h *OrderHandler) DeleteOrderByIDV1(c *gin.Context) {

	id, err := strconv.ParseInt(
		c.Param("id"),
		10,
		64,
	)

	if err != nil {

		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Code:    "INVALID_ORDER_ID",
			Message: "Invalid order id",
		})

		return
	}

	ctx, cancel := context.WithTimeout(
		c.Request.Context(),
		2*time.Second,
	)
	defer cancel()

	err = h.orderUsecase.DeleteOrderByID(
		ctx,
		id,
	)

	if err != nil {

		if errors.Is(err, model.ErrOrderNotFound) {

			c.JSON(http.StatusNotFound, model.ErrorResponse{
				Code:    "ORDER_NOT_FOUND",
				Message: "Order not found",
			})

			return
		}

		log.Println("Error:", err)

		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Code:    "INTERNAL_SERVER_ERROR",
			Message: "Internal server error",
		})

		return
	}

	c.JSON(http.StatusOK, model.DeleteOrderResponse{
		ID:      id,
		Message: "Deleted order successfully",
	})
}
