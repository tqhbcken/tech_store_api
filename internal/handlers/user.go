package handlers

import (
	"api_techstore/internal/models"
	"api_techstore/internal/services"
	"api_techstore/pkg/response"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	// "golang.org/x/crypto/bcrypt"
)

type UserReq struct {
	Id string `json:"id" binding:"required"`
	models.UserReq
}

func GetAllUsers(c *gin.Context) {
	//goi toi service
	users, err := services.GetAllUsers()

	//neu co loi thi tra ve loi
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	//neu khong co loi thi tra ve ds users
	response.SuccessResponse(c, http.StatusOK, "Users retrieved successfully", users)
}

func GetUserById(c *gin.Context) {
	//lay id tu client
	id := c.Param("id")

	//goi toi service
	user, err := services.GetUserById(id)

	//neu co loi thi tra ve loi
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	//neu khong co loi thi tra ve user
	response.SuccessResponse(c, http.StatusOK, "User retrieved successfully", user)
}

func CreateUser(c *gin.Context) {
	//lay data tu body
	var userReq models.UserReq
	if err := c.ShouldBindJSON(&userReq); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// Hash password
	hashedPassword, error := bcrypt.GenerateFromPassword([]byte(userReq.Password), bcrypt.DefaultCost)
	if error != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, "Failed to hash password")
		return
	}

	// Convert UserReq to User
	user := models.User{
		FullName:      userReq.FullName,
		Email:         userReq.Email,
		Phone:         userReq.Phone,
		PasswordHash:  string(hashedPassword),
		Role:          userReq.Role,
		IsActive:      userReq.IsActive,
	}

	//goi toi service
	err := services.CreateUser(user)

	//neu co loi thi tra ve loi
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	//neu khong co loi thi tra ve thanh cong
	response.SuccessResponse(c, http.StatusCreated, "User created successfully", nil)
}

func UpdateUser(c *gin.Context) {
	//lay id tu client
	id := c.Param("id")

	//kiem tra id
	if id == "" {	
		response.ErrorResponse(c, http.StatusBadRequest, "ID is required")
		return
	}

	//goi toi service de kiem tra user co ton tai khong
	idCheker, error := services.GetUserById(id)
	if idCheker.ID == 0 || error != nil {
		response.ErrorResponse(c, http.StatusNotFound, "User not found")
		return
	}

	//lay data tu body
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	//goi toi service
	err := services.UpdateUser(id, user)

	//neu co loi thi tra ve loi
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	//neu khong co loi thi tra ve thanh cong
	response.SuccessResponse(c, http.StatusOK, "User updated successfully", nil)
}

func DeleteUser(c *gin.Context) {
	//lay id tu param
	id := c.Param("id")

	//goi toi service de kiem tra user co ton tai khong
	idCheker, error := services.GetUserById(id)
	if idCheker.ID == 0 || error != nil {
		response.ErrorResponse(c, http.StatusNotFound, "User not found")
		return
	}

	//goi toi service
	err := services.DeleteUser(id)

	//neu co loi thi tra ve loi
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	//neu khong co loi thi tra ve thanh cong
	response.SuccessResponse(c, http.StatusNoContent, "User deleted successfully", nil)
}

func GetUserProfile(c *gin.Context) {

	userId, exists := c.Get("user_id")
	if !exists {
		response.ErrorResponse(c, http.StatusUnauthorized, "User ID not found in context")
		return
	}


	/// Convert userId to string
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

	user, err := services.GetUserById(id)
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve user profile: "+err.Error())
		return
	}

	response.SuccessResponse(c, http.StatusOK, "User profile retrieved successfully", user)

}
