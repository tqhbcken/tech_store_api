package v1

import (
	"api_techstore/internal/container"
	"api_techstore/internal/handlers"
	"api_techstore/internal/middlewares"
	"api_techstore/internal/models"

	"github.com/gin-gonic/gin"
)

func SetupPaymentRoutes(r *gin.RouterGroup, public *gin.RouterGroup, ctn *container.Container) {
	// Protected routes (require auth)
	payments := r.Group("/payments")
	{
		payments.POST("",
			middlewares.RequireRole("user", "admin"),
			middlewares.ValidateRequest(&models.PaymentCreateRequest{}),
			func(ctx *gin.Context) {
				handlers.CreatePayment(ctx, ctn)
			})
		payments.GET("/:orderId/status", func(ctx *gin.Context) {
			handlers.GetPaymentStatus(ctx, ctn)
		})
	}

	// Public routes (for callbacks from payment gateways)
	paymentCallbacks := public.Group("/payments-callback")
	{
		// The actual path will depend on the payment provider, e.g., /momo, /vnpay
		paymentCallbacks.POST("/notify", func(ctx *gin.Context) {
			handlers.HandlePaymentCallback(ctx, ctn)
		})
	}
}
