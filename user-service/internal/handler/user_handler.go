package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"ecommerce-microservice/user-service/internal/model"
	"ecommerce-microservice/user-service/internal/usecase"
)

type UserHandler struct {
	userUsecase *usecase.UserUsecase
}

func NewUserHandler(userUsecase *usecase.UserUsecase) *UserHandler {
	return &UserHandler{
		userUsecase: userUsecase,
	}
}

// Versioning For Old Client V1 (http://127.0.0.1:8080/api/v1/users)
func (h *UserHandler) CreateUserV1(c *gin.Context) {
	var request model.CreateUserRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user, err := h.userUsecase.CreateUser(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	response := model.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}

	c.JSON(http.StatusCreated, response)

}

func (h *UserHandler) GetUserByIDV1(c *gin.Context) {
	id := c.Param("id")
	id1, _ := strconv.Atoi(id)
	user, err := h.userUsecase.GetUserByID(int64(id1))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "user not found",
		})
		return
	}

	response := model.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}

	c.JSON(http.StatusOK, response)

}

// Versioning For New Client V2 (http://127.0.0.1:8080/api/v2/users)
func (h *UserHandler) CreateUserV2(c *gin.Context) {
	var request model.CreateUserRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user, err := h.userUsecase.CreateUser(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// response := model.UserResponse{
	// 	ID:    user.ID,
	// 	Name:  user.Name,
	// 	Email: user.Email,
	// 	CreatedAt: user.CreatedAt,
	// }

	type Response struct {
		ID    int64
		Name  string
		Email string
	}
	response := Response{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	c.JSON(http.StatusCreated, response)

}

func (h *UserHandler) GetUserByIDV2(c *gin.Context) {
	id := c.Param("id")
	id1, _ := strconv.Atoi(id)
	user, err := h.userUsecase.GetUserByID(int64(id1))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "user not found",
		})
		return
	}

	// response := model.UserResponse{
	// 	ID:    user.ID,
	// 	Name:  user.Name,
	// 	Email: user.Email,
	// 	CreatedAt: user.CreatedAt,
	// }
	type Response struct {
		ID    int64
		Name  string
		Email string
	}
	response := Response{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	c.JSON(http.StatusOK, response)

}
