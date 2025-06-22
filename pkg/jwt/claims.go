package jwt

import (
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID   uint   `json:"user_id"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

type AccessTokenClaims struct {
	UserID   uint   `json:"user_id"`
	Role     string `json:"role"`
	AccessUUID string `json:"access_uuid"`
	jwt.RegisteredClaims
}

type RefreshTokenClaims struct {
	UserID   uint   `json:"user_id"`
	Role     string `json:"role"`
	RefreshUUID string `json:"refresh_uuid"`
	jwt.RegisteredClaims
}