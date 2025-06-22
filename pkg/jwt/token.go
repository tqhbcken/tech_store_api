package jwt

// import (
// 	"errors"
// 	"time"

// 	"github.com/golang-jwt/jwt/v5"
// 	// "github.com/google/uuid"
// )

// var (
// 	ErrInvalidToken = errors.New("invalid token")
// 	ErrExpiredToken = errors.New("token has expired")
// )


// /// create jwt token with basic auth
// func (c *JWTConfig) GenerateToken(userID uint, role string) (string, error) {
// 	claims := &Claims{
// 		UserID: userID,
// 		Role:   role,
// 		RegisteredClaims: jwt.RegisteredClaims{
// 			ExpiresAt: jwt.NewNumericDate(time.Now().Add(c.TokenDuration)),
// 			IssuedAt:  jwt.NewNumericDate(time.Now()),
// 		},
// 	}

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	return token.SignedString([]byte(c.SecretKey))
// }

// func (c *JWTConfig) ValidateToken(tokenString string) (*Claims, error) {
// 	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
// 		return []byte(c.SecretKey), nil
// 	})

// 	if err != nil {
// 		if errors.Is(err, jwt.ErrTokenExpired) {
// 			return nil, ErrExpiredToken
// 		}
// 		return nil, ErrInvalidToken
// 	}

// 	claims, ok := token.Claims.(*Claims)
// 	if !ok || !token.Valid {
// 		return nil, ErrInvalidToken
// 	}

// 	return claims, nil
// }


