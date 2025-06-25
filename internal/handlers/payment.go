package handlers

import (
	"api_techstore/internal/container"
	"api_techstore/internal/models"
	"api_techstore/pkg/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreatePayment creates a new payment record
func CreatePayment(c *gin.Context, ctn *container.Container) {
	var req models.Payment
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, "Invalid input")
		return
	}
	payment, err := ctn.PaymentService.CreatePayment(req)
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.SuccessResponse(c, http.StatusCreated, "Payment created successfully", payment)
}

// GetPaymentStatus retrieves the status of a payment
func GetPaymentStatus(c *gin.Context, ctn *container.Container) {
	orderIDStr := c.Param("orderId")
	orderID, err := strconv.ParseUint(orderIDStr, 10, 64)
	if err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, "Invalid order id")
		return
	}
	payment, err := ctn.PaymentService.GetPaymentByOrderID(uint(orderID))
	if err != nil {
		response.ErrorResponse(c, http.StatusNotFound, "Payment not found")
		return
	}
	response.SuccessResponse(c, http.StatusOK, "Payment status retrieved successfully", payment.Status)
}

// HandlePaymentCallback handles the callback from a payment gateway (e.g., Momo, VNPay)
func HandlePaymentCallback(c *gin.Context, ctn *container.Container) {
	orderIDStr := c.Query("order_id")
	status := c.Query("status")
	orderID, err := strconv.ParseUint(orderIDStr, 10, 64)
	if err != nil || status == "" {
		response.ErrorResponse(c, http.StatusBadRequest, "Invalid callback data")
		return
	}
	if err := ctn.PaymentService.UpdatePaymentStatus(uint(orderID), status); err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.SuccessResponse(c, http.StatusOK, "Payment status updated successfully", nil)
}
