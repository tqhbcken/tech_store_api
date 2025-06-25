package handlers

import (
	"api_techstore/internal/container"
	"api_techstore/internal/models"
	"api_techstore/pkg/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetCart retrieves the current user's cart
func GetCart(c *gin.Context, ctn *container.Container) {
	userID, exists := c.Get("user_id")
	if !exists {
		response.ErrorResponse(c, http.StatusUnauthorized, "Unauthorized")
		return
	}
	uid, ok := userID.(uint)
	if !ok {
		response.ErrorResponse(c, http.StatusInternalServerError, "Invalid user id type")
		return
	}
	cart, err := ctn.CartService.GetCartByUserID(uid)
	if err != nil {
		response.ErrorResponse(c, http.StatusNotFound, "Cart not found")
		return
	}
	response.SuccessResponse(c, http.StatusOK, "Cart retrieved successfully", cart)
}

// AddItemToCart adds a product to the cart
func AddItemToCart(c *gin.Context, ctn *container.Container) {
	var req models.CartItem
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, "Invalid input")
		return
	}
	item, err := ctn.CartItemService.AddItemToCart(req)
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.SuccessResponse(c, http.StatusCreated, "Item added to cart", item)
}

// UpdateCartItem updates the quantity of an item in the cart
func UpdateCartItem(c *gin.Context, ctn *container.Container) {
	itemIDStr := c.Param("itemId")
	itemID, err := strconv.ParseUint(itemIDStr, 10, 64)
	if err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, "Invalid item id")
		return
	}
	var req models.CartItem
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, "Invalid input")
		return
	}
	item, err := ctn.CartItemService.UpdateCartItem(uint(itemID), req)
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.SuccessResponse(c, http.StatusOK, "Cart item updated", item)
}

// RemoveItemFromCart removes an item from the cart
func RemoveItemFromCart(c *gin.Context, ctn *container.Container) {
	itemIDStr := c.Param("itemId")
	itemID, err := strconv.ParseUint(itemIDStr, 10, 64)
	if err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, "Invalid item id")
		return
	}
	if err := ctn.CartItemService.RemoveItemFromCart(uint(itemID)); err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.SuccessResponse(c, http.StatusOK, "Item removed from cart", nil)
}

// ClearCart removes all items from the cart
func ClearCart(c *gin.Context, ctn *container.Container) {
	cartIDStr := c.Query("cart_id")
	cartID, err := strconv.ParseUint(cartIDStr, 10, 64)
	if err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, "Invalid cart id")
		return
	}
	if err := ctn.CartItemService.ClearCart(uint(cartID)); err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.SuccessResponse(c, http.StatusOK, "Cart cleared", nil)
}
