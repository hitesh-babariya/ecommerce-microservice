package handler

import (
	"context"
	"net/http"
	"strconv"
	"time"

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
	// 2-second db timeout / Or c.Request.Context() you can pass directly
	ctx, cancel := context.WithTimeout(
		c.Request.Context(),
		2*time.Second,
	)
	defer cancel()

	user, err := h.userUsecase.CreateUser(ctx, request)
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

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid user id",
		})
		return
	}

	// 2-second db timeout / Or c.Request.Context() you can pass directly
	ctx, cancel := context.WithTimeout(
		c.Request.Context(),
		2*time.Second,
	)
	defer cancel()

	user, err := h.userUsecase.GetUserByID(ctx, id)
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
func (h *UserHandler) UpdateUserByIDV1(c *gin.Context) {
	id := c.Param("id")
	//id1, err := strconv.Atoi(id) // THis will convert into int not int64
	id1, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid user id",
		})
		return
	}

	var req model.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Invalid Request",
		})
		return
	}

	// 2-second db timeout / Or c.Request.Context() you can pass directly
	ctx, cancel := context.WithTimeout(
		c.Request.Context(),
		2*time.Second,
	)
	defer cancel()

	user, err := h.userUsecase.UpdateUserByID(ctx, req, id1)
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

// Using PATCH Partailly update user with ID
func (h *UserHandler) PatchUserByIDV1(c *gin.Context) {

	id1, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid user id",
		})
		return
	}

	var req model.PatchUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Invalid Request",
		})
		return
	}

	// 2-second db timeout / Or c.Request.Context() you can pass directly
	ctx, cancel := context.WithTimeout(
		c.Request.Context(),
		2*time.Second,
	)
	defer cancel()

	user, err := h.userUsecase.PatchUserByID(ctx, req, id1)
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

// Using DELETE delete user with ID
func (h *UserHandler) DeleteUserByIDV1(c *gin.Context) {

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid user id",
		})
		return
	}

	// 2-second db timeout / Or c.Request.Context() you can pass directly
	ctx, cancel := context.WithTimeout(
		c.Request.Context(),
		2*time.Second,
	)
	defer cancel()

	err = h.userUsecase.DeleteUserByID(ctx, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not found",
		})
		return
	}

	type DeleteUser struct {
		ID      int64
		Message string
	}
	response := DeleteUser{
		ID:      id,
		Message: "Deleted user successfully",
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

	// 2-second db timeout / Or c.Request.Context() you can pass directly
	ctx, cancel := context.WithTimeout(
		c.Request.Context(),
		2*time.Second,
	)
	defer cancel()
	user, err := h.userUsecase.CreateUser(ctx, request)
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
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	// 2-second db timeout / Or c.Request.Context() you can pass directly
	ctx, cancel := context.WithTimeout(
		c.Request.Context(),
		2*time.Second,
	)
	defer cancel()

	user, err := h.userUsecase.GetUserByID(ctx, id)
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

func (h *UserHandler) UpdateUserByIDV2(c *gin.Context) {
	id := c.Param("id")
	id1, _ := strconv.ParseInt(id, 10, 64)

	var req model.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Invalid Request",
		})
		return
	}

	// 2-second db timeout / Or c.Request.Context() you can pass directly
	ctx, cancel := context.WithTimeout(
		c.Request.Context(),
		2*time.Second,
	)
	defer cancel()

	user, err := h.userUsecase.UpdateUserByID(ctx, req, id1)
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

// Using PATCH Partailly update user with ID
func (h *UserHandler) PatchUserByIDV2(c *gin.Context) {

	id1, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid user id",
		})
		return
	}

	var req model.PatchUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Invalid Request",
		})
		return
	}

	// 2-second db timeout / Or c.Request.Context() you can pass directly
	ctx, cancel := context.WithTimeout(
		c.Request.Context(),
		2*time.Second,
	)
	defer cancel()

	user, err := h.userUsecase.PatchUserByID(ctx, req, id1)
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
