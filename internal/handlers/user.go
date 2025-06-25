package handlers

import (
	"api_techstore/internal/container"
	"api_techstore/internal/models"
	"api_techstore/pkg/response"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func GetAllUsers(c *gin.Context, ctn *container.Container) {
	users, err := ctn.UserService.GetAllUsers()
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if len(users) == 0 {
		response.SuccessResponse(c, http.StatusOK, "No users found", nil)
		return
	}
	response.SuccessResponse(c, http.StatusOK, "Users retrieved successfully", users)
}

func GetUserById(c *gin.Context, ctn *container.Container) {
	id := c.Param("id")
	user, err := ctn.UserService.GetUserById(id)
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if user.ID == 0 {
		response.ErrorResponse(c, http.StatusNotFound, "User not found")
		return
	}
	response.SuccessResponse(c, http.StatusOK, "User retrieved successfully", user)
}

func CreateUser(c *gin.Context, ctn *container.Container) {
	var userReq models.CreateUserReq
	if err := c.ShouldBindJSON(&userReq); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if userReq.FullName == "" || userReq.Email == "" || userReq.Phone == "" || userReq.Password == "" {
		response.ErrorResponse(c, http.StatusBadRequest, "FullName, Email, Phone, and Password are required fields")
		return
	}
	// TODO: Implement GetUserByEmail in UserService if needed
	// existingUsers, err := ctn.UserService.GetUserByEmail(userReq.Email)
	// if err != nil {
	// 	response.ErrorResponse(c, http.StatusInternalServerError, "Failed to check existing user: "+err.Error())
	// 	return
	// }
	// if len(existingUsers) > 0 {
	// 	response.ErrorResponse(c, http.StatusConflict, "Email already exists")
	// 	return
	// }
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userReq.Password), bcrypt.DefaultCost)
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, "Failed to hash password")
		return
	}
	user := models.User{
		FullName:     userReq.FullName,
		Email:        userReq.Email,
		Phone:        userReq.Phone,
		PasswordHash: string(hashedPassword),
		Role:         userReq.Role,
		IsActive:     userReq.IsActive,
	}
	err = ctn.UserService.CreateUser(user)
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.SuccessResponse(c, http.StatusCreated, "User created successfully", nil)
}

func UpdateUser(c *gin.Context, ctn *container.Container) {
	id := c.Param("id")
	if id == "" {
		response.ErrorResponse(c, http.StatusBadRequest, "ID is required")
		return
	}
	idChecker, err := ctn.UserService.GetUserById(id)
	if idChecker.ID == 0 || err != nil {
		response.ErrorResponse(c, http.StatusNotFound, "User not found")
		return
	}
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err = ctn.UserService.UpdateUser(id, user)
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.SuccessResponse(c, http.StatusOK, "User updated successfully", nil)
}

func DeleteUser(c *gin.Context, ctn *container.Container) {
	id := c.Param("id")
	idChecker, err := ctn.UserService.GetUserById(id)
	if idChecker.ID == 0 || err != nil {
		response.ErrorResponse(c, http.StatusNotFound, "User not found")
		return
	}
	err = ctn.UserService.DeleteUser(id)
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.SuccessResponse(c, http.StatusOK, "User deleted successfully", nil)
}

func GetUserProfile(c *gin.Context, ctn *container.Container) {
	userId, exists := c.Get("user_id")
	if !exists {
		response.ErrorResponse(c, http.StatusUnauthorized, "User ID not found in context")
		return
	}
	var id string
	switch v := userId.(type) {
	case string:
		id = v
	case uint:
		id = fmt.Sprintf("%d", v)
	case int:
		id = fmt.Sprintf("%d", v)
	default:
		response.ErrorResponse(c, http.StatusInternalServerError, "Invalid user ID type")
		return
	}
	user, err := ctn.UserService.GetUserById(id)
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve user profile: "+err.Error())
		return
	}
	response.SuccessResponse(c, http.StatusOK, "User profile retrieved successfully", user)
}
