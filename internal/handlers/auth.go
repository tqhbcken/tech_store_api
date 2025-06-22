package handlers

import (
	"api_techstore/internal/models"
	"api_techstore/internal/services"
	"api_techstore/pkg/jwt"
	"api_techstore/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	// Lấy dữ liệu từ request body
	var loginReq models.LoginReq
	if err := c.ShouldBindJSON(&loginReq); err != nil || loginReq.Email == "" || loginReq.Password == "" {
		response.ErrorResponse(c, http.StatusBadRequest, "Invalid request body or missing email/password")
		return
	}
	// Kiểm tra email có tồn tại trên hệ thống không
	users, err := services.GetUserByEmail(loginReq.Email)
	if err != nil || len(users) == 0 {
		response.ErrorResponse(c, http.StatusUnauthorized, "Email does not exist")
		return
	}
	user := users[0]
	// Kiểm tra password so với password đã hash
	if !services.CheckPasswordHash(loginReq.Password, user.PasswordHash) {
		response.ErrorResponse(c, http.StatusUnauthorized, "Incorrect password")
		return
	}
	////////// Tạo token //////////
	// token, err := jwt.NewJWTConfig().GenerateToken(user.ID, user.Role)
	// if err != nil {
	// 	response.ErrorResponse(c, http.StatusInternalServerError, "Failed to generate token: "+err.Error())
	// 	return
	// }

	access_token, err := jwt.NewJWTConfig().GenerateAccessRedisToken(user.ID, user.Role)
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, "Failed to generate token: "+err.Error())
		return
	}

	refresh_token, err := jwt.NewJWTConfig().GenerateRefreshRedisToken(user.ID, user.Role)
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, "Failed to generate token: "+err.Error())
		return
	}

	response.SuccessResponse(c, http.StatusOK, "Login successful", gin.H{
		"user_id":       user.ID,
		"user_role":     user.Role,
		"access_token":  access_token,
		"refresh_token": refresh_token,
	})
}

func Register(c *gin.Context) {
	var userReq models.LoginReq
	if err := c.ShouldBindJSON(&userReq); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, "Invalid request body: "+err.Error())
		return
	}
	// Validate dữ liệu
	if userReq.Email == "" || userReq.Password == "" {
		response.ErrorResponse(c, http.StatusBadRequest, "Email and password are required")
		return
	}
	// Kiểm tra email đã tồn tại chưa
	users, err := services.GetUserByEmail(userReq.Email)
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, "Error checking email: "+err.Error())
		return
	}
	if len(users) > 0 {
		response.ErrorResponse(c, http.StatusConflict, "Email already exists")
		return
	}

}

func Logout(c *gin.Context) {
	// Handle user logout
	response.SuccessResponse(c, http.StatusOK, "Logout successful", nil)
}
func RefreshToken(c *gin.Context) {
	var req struct {
		RefreshToken string `json:"refresh_token" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, "refresh_token field is required")
		return
	}

	// Validate the refresh token
	jwtCfg := jwt.NewJWTConfig()
	claims, err := jwtCfg.ValidateRefreshRedisToken(req.RefreshToken)
	if err != nil {
		response.ErrorResponse(c, http.StatusUnauthorized, "Invalid or expired refresh token")
		return
	}

	// Generate a new access token
	newAccessToken, err := jwtCfg.GenerateAccessRedisToken(claims.UserID, claims.Role)
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, "Could not generate new access token")
		return
	}

	response.SuccessResponse(c, http.StatusOK, "Token refreshed successfully", gin.H{
		"access_token": newAccessToken,
	})
}
