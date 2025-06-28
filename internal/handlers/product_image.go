package handlers

import (
	"api_techstore/internal/container"
	"api_techstore/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AddProductImage godoc
// @Summary Add product image
// @Description Add an image to a product (Admin only)
// @Tags product-images
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param product_id path string true "Product ID"
// @Param image_url formData string true "Image URL"
// @Param is_main formData boolean false "Is main image"
// @Success 201 {object} response.Response "Product image added successfully"
// @Failure 400 {object} response.Response "Invalid request"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 403 {object} response.Response "Forbidden"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /products/{product_id}/images [post]
func AddProductImage(c *gin.Context, ctn *container.Container) {
	// TODO: Implement product image upload logic
	response.SuccessResponse(c, http.StatusCreated, "Product image added successfully", nil)
}

// DeleteProductImage godoc
// @Summary Delete product image
// @Description Delete a product image (Admin only)
// @Tags product-images
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param product_id path string true "Product ID"
// @Param image_id path string true "Image ID"
// @Success 200 {object} response.Response "Product image deleted successfully"
// @Failure 400 {object} response.Response "Invalid request"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 403 {object} response.Response "Forbidden"
// @Failure 404 {object} response.Response "Product image not found"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /products/{product_id}/images/{image_id} [delete]
func DeleteProductImage(c *gin.Context, ctn *container.Container) {
	// TODO: Implement product image deletion logic
	response.SuccessResponse(c, http.StatusOK, "Product image deleted successfully", nil)
}
