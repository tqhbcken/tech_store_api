package handlers

import (
	"api_techstore/internal/container"
	"api_techstore/internal/middlewares"
	"api_techstore/internal/models"
	apperrors "api_techstore/pkg/errors"
	"api_techstore/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetAllOrders godoc
// @Summary Get all orders
// @Description Retrieve all orders
// @Tags orders
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} response.Response{data=[]models.SwaggerOrder} "Orders retrieved successfully"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /order [get]
func GetAllOrders(c *gin.Context, ctn *container.Container) {
	orders, err := ctn.OrderService.GetAllOrders()
	if err != nil {
		response.DatabaseErrorResponse(c, err)
		return
	}
	response.SuccessResponse(c, http.StatusOK, "Orders retrieved successfully", orders)
}

// GetOrderByID godoc
// @Summary Get order by ID
// @Description Retrieve a specific order by ID
// @Tags orders
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Order ID"
// @Success 200 {object} response.Response{data=models.SwaggerOrder} "Order retrieved successfully"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /order/{id} [get]
func GetOrderByID(c *gin.Context, ctn *container.Container) {
	id := c.Param("id")
	order, err := ctn.OrderService.GetOrderByID(id)
	if err != nil {
		response.DatabaseErrorResponse(c, err)
		return
	}
	response.SuccessResponse(c, http.StatusOK, "Order retrieved successfully", order)
}

// CreateOrder godoc
// @Summary Create new order
// @Description Create a new order (User/Admin only)
// @Tags orders
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body models.OrderCreateRequest true "Order data"
// @Success 201 {object} response.Response{data=models.SwaggerOrder} "Order created successfully"
// @Failure 400 {object} response.Response "Invalid request"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 403 {object} response.Response "Forbidden"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /order [post]
func CreateOrder(c *gin.Context, ctn *container.Container) {
	// Lấy user_id từ middleware (context)
	userID, exists := c.Get("user_id")
	if !exists {
		response.NewErrorResponse(c, apperrors.NewUnauthorized())
		return
	}
	userIDUint, ok := userID.(uint)
	if !ok {
		response.HandleError(c, apperrors.New(apperrors.ErrCodeInvalidInput, "Invalid user id type", http.StatusInternalServerError))
		return
	}

	// Lấy validated model (không còn user_id)
	req := middlewares.GetValidatedModel(c).(*models.OrderCreateRequest)

	// get user's cart
	cart, err := ctn.CartService.GetCartByUserID(userIDUint)
	if err != nil {
		response.DatabaseErrorResponse(c, err)
		return
	}

	// Lấy các item trong cart (đã preload Product)
	items, err := ctn.CartItemService.GetItemsByCartID(cart.ID)
	if err != nil {
		response.DatabaseErrorResponse(c, err)
		return
	}

	totalAmount := 0.0
	for _, item := range items {
		if item.Product.ID == 0 {
			response.ErrorResponse(c, http.StatusBadRequest, "Product not found for cart item")
			return
		}
		totalAmount += float64(item.Quantity) * item.Product.Price
	}

	order := models.Order{
		UserID:            userIDUint,
		TotalAmount:       totalAmount,
		ShippingAddressID: req.ShippingAddressID,
	}

	// Set default status if not provided
	if order.Status == "" {
		order.Status = "pending"
	}

	newOrder, err := ctn.OrderService.CreateOrder(order)
	if err != nil {
		response.DatabaseErrorResponse(c, err)
		return
	}

	// Tạo order_items từ cart_items
	var orderItems []models.OrderItem
	for _, item := range items {
		orderItems = append(orderItems, models.OrderItem{
			OrderID:   newOrder.ID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			UnitPrice: item.Product.Price,
		})
	}
	if len(orderItems) > 0 {
		db := ctn.DB
		if err := db.Create(&orderItems).Error; err != nil {
			response.DatabaseErrorResponse(c, err)
			return
		}
	}

	// Reload order để trả về kèm order_items
	var orderWithItems models.Order
	db := ctn.DB
	if err := db.Preload("OrderItems").First(&orderWithItems, newOrder.ID).Error; err != nil {
		response.DatabaseErrorResponse(c, err)
		return
	}

	response.SuccessResponse(c, http.StatusCreated, "Order created successfully", orderWithItems)
}

// UpdateOrder godoc
// @Summary Update order
// @Description Update order information (User/Admin only)
// @Tags orders
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Order ID"
// @Param request body models.OrderUpdateRequest true "Order update data"
// @Success 200 {object} response.Response{data=models.SwaggerOrder} "Order updated successfully"
// @Failure 400 {object} response.Response "Invalid request"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 403 {object} response.Response "Forbidden"
// @Failure 404 {object} response.Response "Order not found"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /order/{id} [put]
func UpdateOrder(c *gin.Context, ctn *container.Container) {
	id := c.Param("id")

	// Check if order exists
	_, err := ctn.OrderService.GetOrderByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			response.NotFoundResponse(c, "Order")
			return
		}
		response.DatabaseErrorResponse(c, err)
		return
	}

	req := middlewares.GetValidatedModel(c).(*models.OrderUpdateRequest)

	order := models.Order{
		UserID:            req.UserID,
		TotalAmount:       req.TotalAmount,
		Status:            req.Status,
		ShippingAddressID: req.ShippingAddressID,
	}

	updatedOrder, err := ctn.OrderService.UpdateOrder(id, order)
	if err != nil {
		response.DatabaseErrorResponse(c, err)
		return
	}

	response.SuccessResponse(c, http.StatusOK, "Order updated successfully", updatedOrder)
}

// DeleteOrder godoc
// @Summary Delete order
// @Description Delete an order (User/Admin only)
// @Tags orders
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Order ID"
// @Success 200 {object} response.Response "Order deleted successfully"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 403 {object} response.Response "Forbidden"
// @Failure 404 {object} response.Response "Order not found"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /order/{id} [delete]
func DeleteOrder(c *gin.Context, ctn *container.Container) {
	id := c.Param("id")
	err := ctn.OrderService.DeleteOrder(id)
	if err != nil {
		response.DatabaseErrorResponse(c, err)
		return
	}

	response.SuccessResponse(c, http.StatusOK, "Order deleted successfully", nil)
}

// GetOrdersByUserID godoc
// @Summary Get orders by user ID
// @Description Retrieve all orders for a specific user
// @Tags orders
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param userId path string true "User ID"
// @Success 200 {object} response.Response{data=[]models.SwaggerOrder} "Orders retrieved successfully"
// @Failure 404 {object} response.Response "No orders found for this user"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /order/user/{userId} [get]
func GetOrdersByUserID(c *gin.Context, ctn *container.Container) {
	userID := c.Param("userId")
	orders, err := ctn.OrderService.GetOrdersByUserID(userID)
	if err != nil {
		response.DatabaseErrorResponse(c, err)
		return
	}

	if len(orders) == 0 {
		response.NotFoundResponse(c, "No orders found for this user")
		return
	}

	response.SuccessResponse(c, http.StatusOK, "Orders retrieved successfully", orders)
}
