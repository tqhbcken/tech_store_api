package handlers

import (
	"api_techstore/internal/cache"
	"api_techstore/internal/container"
	"api_techstore/internal/database"
	"api_techstore/internal/middlewares"
	"api_techstore/internal/models"
	"api_techstore/internal/services"
	"api_techstore/pkg/jwt"
	"api_techstore/pkg/response"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Login godoc
// @Summary User login
// @Description Authenticate user and return JWT tokens
// @Tags auth
// @Accept json
// @Produce json
// @Param request body models.LoginReq true "Login credentials"
// @Success 200 {object} response.Response{data=map[string]interface{}} "Login successful"
// @Failure 400 {object} response.Response "Invalid request"
// @Failure 401 {object} response.Response "Invalid credentials"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /auth/login [post]
func Login(c *gin.Context, ctn *container.Container) {
	// Lấy validated model từ middleware
	req := middlewares.GetValidatedModel(c).(*models.LoginReq)
	// Kiểm tra người dùng
	user, err := services.Login(ctn.DB, req.Email, req.Password)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			response.ErrorResponse(c, http.StatusUnauthorized, "Invalid email or password")
			return
		}
		response.ErrorResponse(c, http.StatusInternalServerError, "Error checking user credentials: "+err.Error())
		return
	}

	// Create new JWT Config
	jwtCfg := ctn.JWTConfig

	// Generate Tokens
	accessToken, err := jwtCfg.GenerateAccessRedisToken(user.ID, user.Role)
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, "Failed to generate access token: "+err.Error())
		return
	}

	refreshToken, err := jwtCfg.GenerateRefreshRedisToken(user.ID, user.Role)
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, "Failed to generate refresh token: "+err.Error())
		return
	}

	// Parse tokens to get UUIDs
	accessClaims, err := jwtCfg.ValidateAccessRedisToken(accessToken)
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, "Failed to parse access token: "+err.Error())
		return
	}

	refreshClaims, err := jwtCfg.ValidateRefreshRedisToken(refreshToken)
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, "Failed to parse refresh token: "+err.Error())
		return
	}

	// Connect to Redis //
	redisClient := cache.NewRedisClient(ctn.Redis)

	// Save tokens to Redis
	err = redisClient.SetToken(c.Request.Context(), accessClaims.AccessUUID, strconv.FormatUint(uint64(user.ID), 10), jwtCfg.AccessTokenDuration)
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, "Failed to save access token to Redis: "+err.Error())
		return
	}

	// Debug log
	// fmt.Printf("DEBUG: Saved access UUID: %s for user: %d\n", accessClaims.AccessUUID, user.ID)

	err = redisClient.SetToken(c.Request.Context(), refreshClaims.RefreshUUID, strconv.FormatUint(uint64(user.ID), 10), jwtCfg.RefreshTokenDuration)
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, "Failed to save refresh token to Redis: "+err.Error())
		return
	}

	// Debug log
	// fmt.Printf("DEBUG: Saved refresh UUID: %s for user: %d\n", refreshClaims.RefreshUUID, user.ID)

	response.SuccessResponse(c, http.StatusOK, "Login successful", gin.H{
		"user_id":       user.ID,
		"user_role":     user.Role,
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

// Register godoc
// @Summary User registration
// @Description Register a new user account
// @Tags auth
// @Accept json
// @Produce json
// @Param request body models.RegisterReq true "User registration data"
// @Success 201 {object} response.Response "User registered successfully"
// @Failure 400 {object} response.Response "Invalid request"
// @Failure 409 {object} response.Response "Email already exists"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /auth/register [post]
func Register(c *gin.Context, di *container.Container) {
	// Lấy validated model từ middleware
	req := middlewares.GetValidatedModel(c).(*models.RegisterReq)
	// Kiểm tra email đã tồn tại chưa
	users, err := services.GetUserByEmail(di.DB, req.Email)
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, "Error checking email: "+err.Error())
		return
	}
	if len(users) > 0 {
		response.ErrorResponse(c, http.StatusConflict, "Email already exists")
		return
	}
	// Hash password
	hashedPassword, err := services.HashPassword(req.Password)
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, "Failed to hash password")
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
	err = services.CreateUser(di.DB, user)
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.SuccessResponse(c, http.StatusCreated, "User registered successfully", nil)
}

// Logout godoc
// @Summary User logout
// @Description Logout user and invalidate tokens
// @Tags auth
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body object{refresh_token=string} true "Refresh token"
// @Success 200 {object} response.Response "Logout successful"
// @Failure 400 {object} response.Response "Invalid request"
// @Failure 401 {object} response.Response "Unauthorized"
// @Router /auth/logout [post]
func Logout(c *gin.Context, di *container.Container) {
	// Extract access token claims from context
	accessUUID, ok := c.Get("access_uuid")
	// fmt.Println("DEBUG: Access UUID from context:", accessUUID)
	if !ok {
		response.ErrorResponse(c, http.StatusBadRequest, "Could not get access token claims")
		return
	}

	// Get refresh token from request
	var req struct {
		RefreshToken string `json:"refresh_token" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, "refresh_token field is required")
		return
	}

	// Validate refresh token to get its UUID
	jwtCfg := di.JWTConfig
	refreshClaims, err := jwtCfg.ValidateRefreshRedisToken(req.RefreshToken)
	if err != nil {
		response.ErrorResponse(c, http.StatusUnauthorized, "Invalid or expired refresh token")
		return
	}

	// Connect to Redis
	redisClient := cache.NewRedisClient(di.Redis)

	// Delete tokens from Redis
	_ = redisClient.DeleteToken(c.Request.Context(), accessUUID.(string))
	if refreshClaims != nil {
		_ = redisClient.DeleteToken(c.Request.Context(), refreshClaims.RefreshUUID)
	}

	response.SuccessResponse(c, http.StatusOK, "Logout successful", nil)
}

// RefreshToken godoc
// @Summary Refresh access token
// @Description Get new access token using refresh token
// @Tags auth
// @Accept json
// @Produce json
// @Param request body object{refresh_token=string} true "Refresh token"
// @Success 200 {object} response.Response{data=map[string]interface{}} "Token refreshed successfully"
// @Failure 400 {object} response.Response "Invalid request"
// @Failure 401 {object} response.Response "Invalid refresh token"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /auth/refresh [post]
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

	// Parse new access token to get UUID
	newAccessClaims, err := jwtCfg.ValidateAccessRedisToken(newAccessToken)
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, "Failed to parse new access token: "+err.Error())
		return
	}

	// Save new access token to Redis
	redisConn, err := database.InitRedis()
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, "Failed to connect to Redis: "+err.Error())
		return
	}
	redisClient := cache.NewRedisClient(redisConn)
	err = redisClient.SetToken(c.Request.Context(), newAccessClaims.AccessUUID, strconv.FormatUint(uint64(claims.UserID), 10), jwtCfg.AccessTokenDuration)
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, "Failed to save new access token to Redis: "+err.Error())
		return
	}

	response.SuccessResponse(c, http.StatusOK, "Token refreshed successfully", gin.H{
		"access_token": newAccessToken,
	})
}

// TestRedis - Handler để test Redis (chỉ dùng cho development)
func TestRedis(c *gin.Context) {
	// Connect to Redis
	redisConn, err := database.InitRedis()
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, "Failed to connect to Redis: "+err.Error())
		return
	}

	// Get all keys
	keys, err := redisConn.Keys(c.Request.Context(), "*").Result()
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, "Failed to get Redis keys: "+err.Error())
		return
	}

	// Get values for access and refresh tokens
	tokenData := make(map[string]interface{})
	tokenTTL := make(map[string]interface{})
	for _, key := range keys {
		if strings.HasPrefix(key, "access-") || strings.HasPrefix(key, "refresh-") {
			value, err := redisConn.Get(c.Request.Context(), key).Result()
			if err == nil {
				tokenData[key] = value
			}

			// Get TTL
			ttl, err := redisConn.TTL(c.Request.Context(), key).Result()
			if err == nil {
				tokenTTL[key] = ttl.String()
			}
		}
	}

	response.SuccessResponse(c, http.StatusOK, "Redis test", gin.H{
		"all_keys":   keys,
		"token_data": tokenData,
		"token_ttl":  tokenTTL,
	})
}

// ClearRedis - Handler để xóa thủ công Redis (chỉ dùng cho development)
func ClearRedis(c *gin.Context) {
	// Connect to Redis
	redisConn, err := database.InitRedis()
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, "Failed to connect to Redis: "+err.Error())
		return
	}

	// Get all keys
	keys, err := redisConn.Keys(c.Request.Context(), "*").Result()
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, "Failed to get Redis keys: "+err.Error())
		return
	}

	// Count tokens before deletion
	accessCount := 0
	refreshCount := 0
	for _, key := range keys {
		if strings.HasPrefix(key, "access-") {
			accessCount++
		}
		if strings.HasPrefix(key, "refresh-") {
			refreshCount++
		}
	}

	// Delete all access and refresh tokens
	deletedCount := 0
	for _, key := range keys {
		if strings.HasPrefix(key, "access-") || strings.HasPrefix(key, "refresh-") {
			err := redisConn.Del(c.Request.Context(), key).Err()
			if err == nil {
				deletedCount++
			}
		}
	}

	response.SuccessResponse(c, http.StatusOK, "Redis cleared", gin.H{
		"deleted_tokens":        deletedCount,
		"access_tokens_before":  accessCount,
		"refresh_tokens_before": refreshCount,
	})
}
