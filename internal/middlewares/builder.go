package middlewares

import (
	// "api_techstore/pkg/jwt"

	"github.com/gin-gonic/gin"
)

type MiddlewareBuilder struct {
	middlewares []gin.HandlerFunc
}

// func (b *MiddlewareBuilder) WithAuth(jwtConfig *jwt.JWTConfig) *MiddlewareBuilder {
// 	b.middlewares = append(b.middlewares, JWTAuthMiddleware())
// 	return b
// }

func (b *MiddlewareBuilder) WithRole(roles ...string) *MiddlewareBuilder {
	b.middlewares = append(b.middlewares, RequireRole(roles...))
	return b
}

// WithValidation adds request body validation middleware
func (b *MiddlewareBuilder) WithValidation(model interface{}) *MiddlewareBuilder {
	b.middlewares = append(b.middlewares, ValidateRequest(model))
	return b
}

// WithQueryValidation adds query parameter validation middleware
func (b *MiddlewareBuilder) WithQueryValidation(model interface{}) *MiddlewareBuilder {
	b.middlewares = append(b.middlewares, ValidateQuery(model))
	return b
}

// WithFormValidation adds form data validation middleware
func (b *MiddlewareBuilder) WithFormValidation(model interface{}) *MiddlewareBuilder {
	b.middlewares = append(b.middlewares, ValidateForm(model))
	return b
}

// WithParamValidation adds URL parameter validation middleware
func (b *MiddlewareBuilder) WithParamValidation(model interface{}) *MiddlewareBuilder {
	b.middlewares = append(b.middlewares, ValidateParams(model))
	return b
}

func (b *MiddlewareBuilder) Build() []gin.HandlerFunc {
	return b.middlewares
}
