package handlers

import (
	"api_techstore/internal/models"
	"api_techstore/internal/services"
	"api_techstore/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllOrders(r *gin.Context) {
	order, err := services.GetAllOrders()
	if err != nil {
		response.ErrorResponse(r, http.StatusInternalServerError, err.Error())
		return
	}
	response.SuccessResponse(r, http.StatusOK, "Orders retrieved successfully", order)
}

func GetOrderByID(r *gin.Context) {
	id := r.Param("id")
	order, err := services.GetOrderByID(id)
	if err != nil {
		response.ErrorResponse(r, http.StatusInternalServerError, err.Error())
		return
	}
	response.SuccessResponse(r, http.StatusOK, "Order retrieved successfully", order)
}

func CreateOrder(r *gin.Context) {
	
	var order models.CreateOrderReq
	if err := r.ShouldBindJSON(&order); err != nil {
		response.ErrorResponse(r, http.StatusBadRequest, "Invalid request body: "+err.Error())
		return
	}

	// Convert CreateOrderReq to Order model
	createdOrder := models.Order{
		UserID:      order.UserID,
		TotalAmount: order.TotalAmount,
		Status:      order.Status,
	}

	createdOrder, err := services.CreateOrder(createdOrder)
	if err != nil {
		response.ErrorResponse(r, http.StatusInternalServerError, err.Error())
		return
	}
	response.SuccessResponse(r, http.StatusCreated, "Order created successfully", createdOrder)
}

func UpdateOrder(r *gin.Context) {
	id := r.Param("id")
	var order models.UpdateOrderReq
	if err := r.ShouldBindJSON(&order); err != nil {
		response.ErrorResponse(r, http.StatusBadRequest, "Invalid request body: "+err.Error())
		return
	}

	// Check if the order exists
	if !CheckOrderExists(id) {
		response.ErrorResponse(r, http.StatusNotFound, "Order not found")
		return
	}

	// Convert UpdateOrderReq to Order model
	orderModel := models.Order{
		UserID:      order.UserID,
		TotalAmount: order.TotalAmount,
		Status:      order.Status,
	}

	updatedOrder, err := services.UpdateOrder(id, orderModel)
	if err != nil {
		response.ErrorResponse(r, http.StatusInternalServerError, err.Error())
		return
	}
	response.SuccessResponse(r, http.StatusOK, "Order updated successfully", updatedOrder)
}

func DeleteOrder(r *gin.Context) {
	id := r.Param("id")

	// Check if the order exists
	if !CheckOrderExists(id) {
		response.ErrorResponse(r, http.StatusNotFound, "Order not found")
		return
	}

	if err := services.DeleteOrder(id); err != nil {
		response.ErrorResponse(r, http.StatusInternalServerError, err.Error())
		return
	}
	response.SuccessResponse(r, http.StatusOK, "Order deleted successfully", nil)
}

func CheckOrderExists(id string) bool {
	_, err := services.GetOrderByID(id)
	if err != nil {
		return false // Order does not exist
	}
	return true // Order exists
}