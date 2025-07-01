package middlewares

import (
	"net/http"
	"strings"

	"api_techstore/internal/cache"
	"api_techstore/internal/container"
	apperrors "api_techstore/pkg/errors"
	jwtpkg "api_techstore/pkg/jwt"
	"api_techstore/pkg/response"

	"github.com/gin-gonic/gin"
)

// check jwt in request header
func JWTAuthMiddleware(ctn *container.Container) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			response.NewErrorResponse(ctx, apperrors.NewUnauthorized())
			ctx.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			appErr := apperrors.New(apperrors.ErrCodeTokenInvalid, "Invalid authorization header format", http.StatusUnauthorized)
			response.NewErrorResponse(ctx, appErr)
			ctx.Abort()
			return
		}

		claims, err := ctn.JWTConfig.ValidateAccessRedisToken(parts[1])
		if err != nil {
			var appErr *apperrors.AppError
			if err == jwtpkg.ErrExpiredToken {
				appErr = apperrors.NewTokenExpired()
			} else {
				appErr = apperrors.NewTokenInvalid()
			}
			response.NewErrorResponse(ctx, appErr)
			ctx.Abort()
			return
		}

		// Check if token exists in Redis
		redisClient := cache.NewRedisClient(ctn.Redis)
		isValid, err := redisClient.IsValidToken(ctx.Request.Context(), claims.AccessUUID)
		if err != nil || !isValid {
			response.NewErrorResponse(ctx, apperrors.NewTokenRevoked())
			ctx.Abort()
			return
		}

		ctx.Set("user_id", claims.UserID)
		ctx.Set("role", claims.Role)
		ctx.Set("access_uuid", claims.AccessUUID)

		ctx.Next()
	}
}

// authorization middleware to check user roles
func RequireRole(roles ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userRole, exists := ctx.Get("role")
		if !exists {
			response.NewErrorResponse(ctx, apperrors.NewUnauthorized())
			ctx.Abort()
			return
		}

		for _, role := range roles {
			if role == userRole {
				ctx.Next()
				return
			}
		}

		response.NewErrorResponse(ctx, apperrors.NewForbidden())
		ctx.Abort()
	}
}
