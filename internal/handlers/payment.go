package handlers

import (
	"api_techstore/internal/container"
	"api_techstore/internal/middlewares"
	"api_techstore/internal/models"
	"api_techstore/pkg/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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
	paymentModel := models.Payment{
		OrderID: req.OrderID,
		Amount:  req.Amount,
		Method:  req.Method,
		Status:  "pending",
	}
	payment, err := ctn.PaymentService.CreatePayment(paymentModel)
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.SuccessResponse(c, http.StatusCreated, "Payment created successfully", payment)
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
