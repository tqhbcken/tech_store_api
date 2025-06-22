package middlewares

import (
	"net/http"
	"strings"

	jwtpkg "api_techstore/pkg/jwt"
	"api_techstore/pkg/response"

	"github.com/gin-gonic/gin"
)

// check jwt in request header
func JWTAuthMiddleware(config *jwtpkg.JWTConfig) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			response.ErrorResponse(
				ctx, http.StatusUnauthorized, "Authorization header is required",
			)
			ctx.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			response.ErrorResponse(
				ctx, http.StatusUnauthorized, "Invalid authorization header format",
			)
			ctx.Abort()
			return
		}

		claims, err := config.ValidateAccessRedisToken(parts[1])
		if err != nil {
			message := "Invalid token"
			if err == jwtpkg.ErrExpiredToken {
				message = "Token has expired"
			}
			response.ErrorResponse(ctx, http.StatusUnauthorized, message)
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
			response.ErrorResponse(ctx, http.StatusForbidden, "Role not found in context")
			ctx.Abort()
			return
		}

		for _, role := range roles {
			if role == userRole {
				ctx.Next()
				return
			}
		}

		response.ErrorResponse(ctx, http.StatusForbidden, "Insufficient permissions")
		ctx.Abort()
	}
}
