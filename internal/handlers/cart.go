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

// GetCart godoc
// @Summary Get cart
// @Description Retrieve the current user's cart
// @Tags cart
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} response.Response{data=models.SwaggerCart} "Cart retrieved successfully"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /cart [get]
func GetCart(c *gin.Context, ctn *container.Container) {
	// Get user ID from context
	userID, exists := c.Get("user_id")
	if !exists {
		response.NewErrorResponse(c, apperrors.NewUnauthorized())
		return
	}

	// Convert user ID to uint
	userIDUint, ok := userID.(uint)
	if !ok {
		response.HandleError(c, apperrors.New(apperrors.ErrCodeInvalidInput, "Invalid user id type", http.StatusInternalServerError))
		return
	}

	// Get or create cart
	cart, err := ctn.CartService.GetCartByUserID(userIDUint)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// Create new cart if not exists
			newCart := models.Cart{
				UserID:    &userIDUint,
				SessionID: "session_" + strconv.FormatUint(uint64(userIDUint), 10),
				Status:    "active",
			}
			cart, err = ctn.CartService.CreateCart(newCart)
			if err != nil {
				response.DatabaseErrorResponse(c, err)
				return
			}
		} else {
			response.DatabaseErrorResponse(c, err)
			return
		}
	}

	// Get cart items
	items, err := ctn.CartItemService.GetItemsByCartID(cart.ID)
	if err != nil {
		response.DatabaseErrorResponse(c, err)
		return
	}

	cart.Items = items
	response.SuccessResponse(c, http.StatusOK, "Cart retrieved successfully", cart)
}

// AddItemToCart godoc
// @Summary Add item to cart
// @Description Add a product to the current user's cart
// @Tags cart
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body models.CartAddItemRequest true "Cart item data"
// @Success 201 {object} response.Response{data=models.SwaggerCartItem} "Item added to cart"
// @Failure 400 {object} response.Response "Invalid request"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /cart/items [post]
func AddItemToCart(c *gin.Context, ctn *container.Container) {
	// Get user ID from context
	userID, exists := c.Get("user_id")
	if !exists {
		response.NewErrorResponse(c, apperrors.NewUnauthorized())
		return
	}

	// Convert user ID to uint
	userIDUint, ok := userID.(uint)
	if !ok {
		response.HandleError(c, apperrors.New(apperrors.ErrCodeInvalidInput, "Invalid user id type", http.StatusInternalServerError))
		return
	}

	// Get or create cart
	cart, err := ctn.CartService.GetCartByUserID(userIDUint)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// Create new cart if not exists
			newCart := models.Cart{
				UserID:    &userIDUint,
				SessionID: "session_" + strconv.FormatUint(uint64(userIDUint), 10),
				Status:    "active",
			}
			cart, err = ctn.CartService.CreateCart(newCart)
			if err != nil {
				response.DatabaseErrorResponse(c, err)
				return
			}
		} else {
			response.DatabaseErrorResponse(c, err)
			return
		}
	}

	// Get validated model from middleware
	req := middlewares.GetValidatedModel(c).(*models.CartAddItemRequest)

	cartItem := models.CartItem{
		CartID:    cart.ID,
		ProductID: req.ProductID,
		Quantity:  req.Quantity,
	}

	newItem, err := ctn.CartItemService.AddItemToCart(cartItem)
	if err != nil {
		response.DatabaseErrorResponse(c, err)
		return
	}

	response.SuccessResponse(c, http.StatusCreated, "Item added to cart", newItem)
}

// UpdateCartItem godoc
// @Summary Update cart item
// @Description Update the quantity of an item in the cart
// @Tags cart
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param itemId path string true "Cart Item ID"
// @Param request body models.CartUpdateItemRequest true "Cart item update data"
// @Success 200 {object} response.Response{data=models.SwaggerCartItem} "Cart item updated"
// @Failure 400 {object} response.Response "Invalid request"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 404 {object} response.Response "Cart item not found"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /cart/items/{itemId} [put]
func UpdateCartItem(c *gin.Context, ctn *container.Container) {
	// Get user ID from context
	userID, exists := c.Get("user_id")
	if !exists {
		response.NewErrorResponse(c, apperrors.NewUnauthorized())
		return
	}

	// Convert user ID to uint
	userIDUint, ok := userID.(uint)
	if !ok {
		response.HandleError(c, apperrors.New(apperrors.ErrCodeInvalidInput, "Invalid user id type", http.StatusInternalServerError))
		return
	}

	// Get item ID from path
	itemIDStr := c.Param("itemId")
	itemID, err := strconv.ParseUint(itemIDStr, 10, 32)
	if err != nil {
		response.NewErrorResponse(c, apperrors.NewValidationFailed("Invalid item id"))
		return
	}

	// Get cart for user
	cart, err := ctn.CartService.GetCartByUserID(userIDUint)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			response.NotFoundResponse(c, "Cart")
			return
		}
		response.DatabaseErrorResponse(c, err)
		return
	}

	// Get cart items to verify ownership
	items, err := ctn.CartItemService.GetItemsByCartID(cart.ID)
	if err != nil {
		response.DatabaseErrorResponse(c, err)
		return
	}

	// Check if item exists in user's cart
	itemExists := false
	for _, item := range items {
		if item.ID == uint(itemID) {
			itemExists = true
			break
		}
	}

	if !itemExists {
		response.NotFoundResponse(c, "Cart item")
		return
	}

	// Get validated model from middleware
	req := middlewares.GetValidatedModel(c).(*models.CartUpdateItemRequest)

	cartItem := models.CartItem{
		CartID:   cart.ID,
		Quantity: req.Quantity,
	}

	updatedItem, err := ctn.CartItemService.UpdateCartItem(uint(itemID), cartItem)
	if err != nil {
		response.DatabaseErrorResponse(c, err)
		return
	}

	response.SuccessResponse(c, http.StatusOK, "Cart item updated", updatedItem)
}

// RemoveItemFromCart godoc
// @Summary Remove item from cart
// @Description Remove an item from the cart
// @Tags cart
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param itemId path string true "Cart Item ID"
// @Success 200 {object} response.Response "Item removed from cart"
// @Failure 400 {object} response.Response "Invalid request"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 404 {object} response.Response "Cart item not found"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /cart/items/{itemId} [delete]
func RemoveItemFromCart(c *gin.Context, ctn *container.Container) {
	// Get user ID from context
	userID, exists := c.Get("user_id")
	if !exists {
		response.NewErrorResponse(c, apperrors.NewUnauthorized())
		return
	}

	// Convert user ID to uint
	userIDUint, ok := userID.(uint)
	if !ok {
		response.HandleError(c, apperrors.New(apperrors.ErrCodeInvalidInput, "Invalid user id type", http.StatusInternalServerError))
		return
	}

	// Get item ID from path
	itemIDStr := c.Param("itemId")
	itemID, err := strconv.ParseUint(itemIDStr, 10, 32)
	if err != nil {
		response.NewErrorResponse(c, apperrors.NewValidationFailed("Invalid item id"))
		return
	}

	// Get cart for user
	cart, err := ctn.CartService.GetCartByUserID(userIDUint)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			response.NotFoundResponse(c, "Cart")
			return
		}
		response.DatabaseErrorResponse(c, err)
		return
	}

	// Get cart items to verify ownership
	items, err := ctn.CartItemService.GetItemsByCartID(cart.ID)
	if err != nil {
		response.DatabaseErrorResponse(c, err)
		return
	}

	// Check if item exists in user's cart
	itemExists := false
	for _, item := range items {
		if item.ID == uint(itemID) {
			itemExists = true
			break
		}
	}

	if !itemExists {
		response.NotFoundResponse(c, "Cart item")
		return
	}

	// Remove item
	err = ctn.CartItemService.RemoveItemFromCart(uint(itemID))
	if err != nil {
		response.DatabaseErrorResponse(c, err)
		return
	}

	response.SuccessResponse(c, http.StatusOK, "Item removed from cart", nil)
}

// ClearCart godoc
// @Summary Clear cart
// @Description Remove all items from the cart
// @Tags cart
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} response.Response "Cart cleared"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 404 {object} response.Response "Cart not found"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /cart [delete]
func ClearCart(c *gin.Context, ctn *container.Container) {
	// Get user ID from context
	userID, exists := c.Get("user_id")
	if !exists {
		response.NewErrorResponse(c, apperrors.NewUnauthorized())
		return
	}

	// Convert user ID to uint
	userIDUint, ok := userID.(uint)
	if !ok {
		response.HandleError(c, apperrors.New(apperrors.ErrCodeInvalidInput, "Invalid user id type", http.StatusInternalServerError))
		return
	}

	// Get cart for user
	cart, err := ctn.CartService.GetCartByUserID(userIDUint)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			response.NotFoundResponse(c, "Cart")
			return
		}
		response.DatabaseErrorResponse(c, err)
		return
	}

	// Clear cart items
	err = ctn.CartItemService.ClearCart(cart.ID)
	if err != nil {
		response.DatabaseErrorResponse(c, err)
		return
	}

	response.SuccessResponse(c, http.StatusOK, "Cart cleared", nil)
}
