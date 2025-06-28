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
		// Nếu chưa có cart, tạo mới cart rỗng
		cartModel := models.Cart{
			UserID: &uid,
			Status: "active",
		}
		cart, err = ctn.CartService.CreateCart(cartModel)
		if err != nil {
			response.ErrorResponse(c, http.StatusInternalServerError, "Failed to create cart: "+err.Error())
			return
		}
	}

	// Lấy items trong cart
	items, err := ctn.CartItemService.GetItemsByCartID(cart.ID)
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, "Failed to get cart items: "+err.Error())
		return
	}

	// Gán items vào cart
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
	// Lấy userID từ context
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

	// Lấy hoặc tạo cart cho user
	cart, err := ctn.CartService.GetCartByUserID(uid)
	if err != nil {
		// Nếu chưa có cart, tạo mới
		cartModel := models.Cart{
			UserID: &uid,
			Status: "active",
		}
		cart, err = ctn.CartService.CreateCart(cartModel)
		if err != nil {
			response.ErrorResponse(c, http.StatusInternalServerError, "Failed to create cart: "+err.Error())
			return
		}
	}

	// Lấy validated request
	req := middlewares.GetValidatedModel(c).(*models.CartAddItemRequest)

	// Tạo cart item với CartID
	itemModel := models.CartItem{
		CartID:    cart.ID,
		ProductID: req.ProductID,
		Quantity:  req.Quantity,
	}

	item, err := ctn.CartItemService.AddItemToCart(itemModel)
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.SuccessResponse(c, http.StatusCreated, "Item added to cart", item)
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
	// Lấy userID từ context
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

	itemIDStr := c.Param("itemId")
	itemID, err := strconv.ParseUint(itemIDStr, 10, 64)
	if err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, "Invalid item id")
		return
	}

	// Kiểm tra quyền sở hữu cart item
	cart, err := ctn.CartService.GetCartByUserID(uid)
	if err != nil {
		response.ErrorResponse(c, http.StatusNotFound, "Cart not found")
		return
	}

	// Kiểm tra item có thuộc cart của user không
	items, err := ctn.CartItemService.GetItemsByCartID(cart.ID)
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, "Failed to get cart items")
		return
	}

	itemExists := false
	for _, item := range items {
		if item.ID == uint(itemID) {
			itemExists = true
			break
		}
	}

	if !itemExists {
		response.ErrorResponse(c, http.StatusNotFound, "Cart item not found")
		return
	}

	req := middlewares.GetValidatedModel(c).(*models.CartUpdateItemRequest)
	itemModel := models.CartItem{
		Quantity: req.Quantity,
	}
	item, err := ctn.CartItemService.UpdateCartItem(uint(itemID), itemModel)
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.SuccessResponse(c, http.StatusOK, "Cart item updated", item)
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
	// Lấy userID từ context
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

	itemIDStr := c.Param("itemId")
	itemID, err := strconv.ParseUint(itemIDStr, 10, 64)
	if err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, "Invalid item id")
		return
	}

	// Kiểm tra quyền sở hữu cart item
	cart, err := ctn.CartService.GetCartByUserID(uid)
	if err != nil {
		response.ErrorResponse(c, http.StatusNotFound, "Cart not found")
		return
	}

	// Kiểm tra item có thuộc cart của user không
	items, err := ctn.CartItemService.GetItemsByCartID(cart.ID)
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, "Failed to get cart items")
		return
	}

	itemExists := false
	for _, item := range items {
		if item.ID == uint(itemID) {
			itemExists = true
			break
		}
	}

	if !itemExists {
		response.ErrorResponse(c, http.StatusNotFound, "Cart item not found")
		return
	}

	if err := ctn.CartItemService.RemoveItemFromCart(uint(itemID)); err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
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
	// Lấy userID từ context
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

	// Lấy cart của user
	cart, err := ctn.CartService.GetCartByUserID(uid)
	if err != nil {
		response.ErrorResponse(c, http.StatusNotFound, "Cart not found")
		return
	}

	if err := ctn.CartItemService.ClearCart(cart.ID); err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.SuccessResponse(c, http.StatusOK, "Cart cleared", nil)
}
