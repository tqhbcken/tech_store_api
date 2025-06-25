package handlers

import (
	"api_techstore/internal/container"
	"api_techstore/internal/middlewares"
	"api_techstore/internal/models"
	"api_techstore/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllOrders(c *gin.Context, ctn *container.Container) {
	orders, err := ctn.OrderService.GetAllOrders()
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.SuccessResponse(c, http.StatusOK, "Orders retrieved successfully", orders)
}

func GetOrderByID(c *gin.Context, ctn *container.Container) {
	id := c.Param("id")
	order, err := ctn.OrderService.GetOrderByID(id)
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.SuccessResponse(c, http.StatusOK, "Order retrieved successfully", order)
}

func CreateOrder(c *gin.Context, ctn *container.Container) {
	req := middlewares.GetValidatedModel(c).(*models.OrderCreateRequest)
	order := models.Order{
		UserID:      req.UserID,
		TotalAmount: req.TotalAmount,
		Status:      req.Status,
	}
	createdOrder, err := ctn.OrderService.CreateOrder(order)
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.SuccessResponse(c, http.StatusCreated, "Order created successfully", createdOrder)
}

func UpdateOrder(c *gin.Context, ctn *container.Container) {
	id := c.Param("id")
	req := middlewares.GetValidatedModel(c).(*models.OrderUpdateRequest)
	if !checkOrderExists(ctn, id) {
		response.ErrorResponse(c, http.StatusNotFound, "Order not found")
		return
	}
	order := models.Order{
		UserID:      req.UserID,
		TotalAmount: req.TotalAmount,
		Status:      req.Status,
	}
	updatedOrder, err := ctn.OrderService.UpdateOrder(id, order)
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.SuccessResponse(c, http.StatusOK, "Order updated successfully", updatedOrder)
}

func DeleteOrder(c *gin.Context, ctn *container.Container) {
	id := c.Param("id")
	if !checkOrderExists(ctn, id) {
		response.ErrorResponse(c, http.StatusNotFound, "Order not found")
		return
	}
	if err := ctn.OrderService.DeleteOrder(id); err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.SuccessResponse(c, http.StatusOK, "Order deleted successfully", nil)
}

func GetOrdersByUserID(c *gin.Context, ctn *container.Container) {
	userID := c.Param("userId")
	orders, err := ctn.OrderService.GetOrdersByUserID(userID)
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if len(orders) == 0 {
		response.ErrorResponse(c, http.StatusNotFound, "No orders found for this user")
		return
	}
	response.SuccessResponse(c, http.StatusOK, "Orders retrieved successfully", orders)
}

func checkOrderExists(ctn *container.Container, id string) bool {
	_, err := ctn.OrderService.GetOrderByID(id)
	return err == nil
}
