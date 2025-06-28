package handlers

import (
	"api_techstore/internal/container"
	"api_techstore/internal/middlewares"
	"api_techstore/internal/models"
	"api_techstore/pkg/response"
	"net/http"
	"strconv"

	apperrors "api_techstore/pkg/errors"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CreatePayment godoc
// @Summary Create payment
// @Description Create a new payment record (User/Admin only)
// @Tags payments
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body models.PaymentCreateRequest true "Payment data"
// @Success 201 {object} response.Response{data=models.SwaggerPayment} "Payment created successfully"
// @Failure 400 {object} response.Response "Invalid request"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 403 {object} response.Response "Forbidden"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /payments [post]
func CreatePayment(c *gin.Context, ctn *container.Container) {
	req := middlewares.GetValidatedModel(c).(*models.PaymentCreateRequest)

	payment := models.Payment{
		OrderID: req.OrderID,
		Amount:  req.Amount,
		Method:  req.Method,
		Status:  "pending", // Default status
	}

	newPayment, err := ctn.PaymentService.CreatePayment(payment)
	if err != nil {
		response.DatabaseErrorResponse(c, err)
		return
	}

	response.SuccessResponse(c, http.StatusCreated, "Payment created successfully", newPayment)
}

// GetPaymentStatus godoc
// @Summary Get payment status
// @Description Retrieve the status of a payment
// @Tags payments
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param orderId path string true "Order ID"
// @Success 200 {object} response.Response{data=string} "Payment status retrieved successfully"
// @Failure 400 {object} response.Response "Invalid order id"
// @Failure 404 {object} response.Response "Payment not found"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /payments/{orderId}/status [get]
func GetPaymentStatus(c *gin.Context, ctn *container.Container) {
	orderIDStr := c.Param("orderId")
	orderID, err := strconv.ParseUint(orderIDStr, 10, 32)
	if err != nil {
		response.NewErrorResponse(c, apperrors.NewValidationFailed("Invalid order id"))
		return
	}

	payment, err := ctn.PaymentService.GetPaymentByOrderID(uint(orderID))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			response.NotFoundResponse(c, "Payment")
			return
		}
		response.DatabaseErrorResponse(c, err)
		return
	}

	response.SuccessResponse(c, http.StatusOK, "Payment status retrieved successfully", payment.Status)
}

// HandlePaymentCallback godoc
// @Summary Handle payment callback
// @Description Handle the callback from a payment gateway (e.g., Momo, VNPay)
// @Tags payments
// @Accept json
// @Produce json
// @Param order_id query string true "Order ID"
// @Param status query string true "Payment status"
// @Success 200 {object} response.Response "Payment status updated successfully"
// @Failure 400 {object} response.Response "Invalid callback data"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /payments-callback/notify [post]
func HandlePaymentCallback(c *gin.Context, ctn *container.Container) {
	orderID := c.Query("order_id")
	status := c.Query("status")

	if orderID == "" || status == "" {
		response.NewErrorResponse(c, apperrors.NewValidationFailed("Invalid callback data"))
		return
	}

	orderIDUint, err := strconv.ParseUint(orderID, 10, 32)
	if err != nil {
		response.NewErrorResponse(c, apperrors.NewValidationFailed("Invalid order id"))
		return
	}

	err = ctn.PaymentService.UpdatePaymentStatus(uint(orderIDUint), status)
	if err != nil {
		response.DatabaseErrorResponse(c, err)
		return
	}

	response.SuccessResponse(c, http.StatusOK, "Payment status updated successfully", nil)
}
