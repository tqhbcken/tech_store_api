package handlers

import (
	"api_techstore/internal/container"
	"api_techstore/internal/middlewares"
	"api_techstore/internal/models"
	"api_techstore/internal/services"
	"api_techstore/pkg/response"
	"net/http"
	"strconv"

	apperrors "api_techstore/pkg/errors"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetAllUsers godoc
// @Summary Get all users
// @Description Retrieve all users (Admin only)
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} response.Response{data=[]models.SwaggerUser} "Users retrieved successfully"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 403 {object} response.Response "Forbidden"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /users [get]
func GetAllUsers(c *gin.Context, ctn *container.Container) {
	users, err := ctn.UserService.GetAllUsers()
	if err != nil {
		response.DatabaseErrorResponse(c, err)
		return
	}
	response.SuccessResponse(c, http.StatusOK, "Users retrieved successfully", users)
}

// GetUserById godoc
// @Summary Get user by ID
// @Description Retrieve a specific user by ID (Admin only)
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "User ID"
// @Success 200 {object} response.Response{data=models.SwaggerUser} "User retrieved successfully"
// @Failure 400 {object} response.Response "Invalid user ID"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 403 {object} response.Response "Forbidden"
// @Failure 404 {object} response.Response "User not found"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /users/{id} [get]
func GetUserById(c *gin.Context, ctn *container.Container) {
	id := c.Param("id")
	user, err := ctn.UserService.GetUserById(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			response.NotFoundResponse(c, "User")
			return
		}
		response.DatabaseErrorResponse(c, err)
		return
	}
	response.SuccessResponse(c, http.StatusOK, "User retrieved successfully", user)
}

// CreateUser godoc
// @Summary Create new user
// @Description Create a new user account (Admin only)
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body models.CreateUserReq true "User data"
// @Success 201 {object} response.Response "User created successfully"
// @Failure 400 {object} response.Response "Invalid request"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 403 {object} response.Response "Forbidden"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /users [post]
func CreateUser(c *gin.Context, ctn *container.Container) {
	req := middlewares.GetValidatedModel(c).(*models.CreateUserReq)

	// Hash password
	hashedPassword, err := services.HashPassword(req.Password)
	if err != nil {
		response.HandleError(c, err)
		return
	}

	user := models.User{
		FullName:     req.FullName,
		Email:        req.Email,
		Phone:        req.Phone,
		PasswordHash: hashedPassword,
		Role:         req.Role,
		IsActive:     req.IsActive,
	}

	err = ctn.UserService.CreateUser(user)
	if err != nil {
		response.DatabaseErrorResponse(c, err)
		return
	}

	response.SuccessResponse(c, http.StatusCreated, "User created successfully", nil)
}

// UpdateUser godoc
// @Summary Update user
// @Description Update user information (Admin only)
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "User ID"
// @Param request body models.UserUpdateRequest true "User update data"
// @Success 200 {object} response.Response "User updated successfully"
// @Failure 400 {object} response.Response "Invalid request"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 403 {object} response.Response "Forbidden"
// @Failure 404 {object} response.Response "User not found"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /users/{id} [put]
func UpdateUser(c *gin.Context, ctn *container.Container) {
	id := c.Param("id")
	if id == "" {
		response.NewErrorResponse(c, apperrors.NewValidationFailed("ID is required"))
		return
	}

	// Check if user exists
	exists, err := ctn.UserService.CheckUserExists(id)
	if err != nil {
		response.DatabaseErrorResponse(c, err)
		return
	}
	if !exists {
		response.NotFoundResponse(c, "User")
		return
	}

	req := middlewares.GetValidatedModel(c).(*models.UserUpdateRequest)

	user := models.User{
		FullName: req.FullName,
		Email:    req.Email,
		Phone:    req.Phone,
		Role:     req.Role,
		IsActive: false,
	}

	if req.Password != "" {
		hashedPassword, err := services.HashPassword(req.Password)
		if err != nil {
			response.HandleError(c, err)
			return
		}
		user.PasswordHash = hashedPassword
	}

	if req.IsActive != nil {
		user.IsActive = *req.IsActive
	}

	err = ctn.UserService.UpdateUser(id, user)
	if err != nil {
		response.DatabaseErrorResponse(c, err)
		return
	}

	response.SuccessResponse(c, http.StatusOK, "User updated successfully", nil)
}

// DeleteUser godoc
// @Summary Delete user
// @Description Delete a user account (Admin only)
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "User ID"
// @Success 200 {object} response.Response "User deleted successfully"
// @Failure 400 {object} response.Response "Invalid user ID"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 403 {object} response.Response "Forbidden"
// @Failure 404 {object} response.Response "User not found"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /users/{id} [delete]
func DeleteUser(c *gin.Context, ctn *container.Container) {
	id := c.Param("id")
	exists, err := ctn.UserService.CheckUserExists(id)
	if err != nil {
		response.DatabaseErrorResponse(c, err)
		return
	}
	if !exists {
		response.NotFoundResponse(c, "User")
		return
	}

	err = ctn.UserService.DeleteUser(id)
	if err != nil {
		response.DatabaseErrorResponse(c, err)
		return
	}

	response.SuccessResponse(c, http.StatusOK, "User deleted successfully", nil)
}

// GetUserProfile godoc
// @Summary Get user profile
// @Description Get current user's profile information
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} response.Response{data=models.SwaggerUser} "User profile retrieved successfully"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /users/profile [get]
func GetUserProfile(c *gin.Context, ctn *container.Container) {
	// Get user ID from context
	userID, exists := c.Get("user_id")
	if !exists {
		response.NewErrorResponse(c, apperrors.NewUnauthorized())
		return
	}

	// Convert user ID to string
	userIDStr, ok := userID.(uint)
	if !ok {
		response.HandleError(c, apperrors.New(apperrors.ErrCodeInvalidInput, "Invalid user ID type", http.StatusInternalServerError))
		return
	}

	user, err := ctn.UserService.GetUserById(strconv.FormatUint(uint64(userIDStr), 10))
	if err != nil {
		response.DatabaseErrorResponse(c, err)
		return
	}

	response.SuccessResponse(c, http.StatusOK, "User profile retrieved successfully", user)
}
