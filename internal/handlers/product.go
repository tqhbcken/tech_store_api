package handlers

import (
	"api_techstore/internal/container"
	"api_techstore/internal/middlewares"
	"api_techstore/internal/models"
	"api_techstore/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAllProducts godoc
// @Summary Get all products
// @Description Retrieve all products
// @Tags products
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} response.Response{data=[]models.SwaggerProduct} "Products retrieved successfully"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /products [get]
func GetAllProducts(c *gin.Context, ctn *container.Container) {
	products, err := ctn.ProductService.GetAllProducts()
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.SuccessResponse(c, http.StatusOK, "Products retrieved successfully", products)
}

// GetProductById godoc
// @Summary Get product by ID
// @Description Retrieve a specific product by ID
// @Tags products
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Product ID"
// @Success 200 {object} response.Response{data=models.SwaggerProduct} "Product retrieved successfully"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /products/{id} [get]
func GetProductById(c *gin.Context, ctn *container.Container) {
	id := c.Param("id")
	product, err := ctn.ProductService.GetProductById(id)
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.SuccessResponse(c, http.StatusOK, "Product retrieved successfully", product)
}

// CreateProduct godoc
// @Summary Create new product
// @Description Create a new product (Admin only)
// @Tags products
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body models.ProductCreateRequest true "Product data"
// @Success 201 {object} response.Response{data=models.SwaggerProduct} "Product created successfully"
// @Failure 400 {object} response.Response "Invalid request"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 403 {object} response.Response "Forbidden"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /products [post]
func CreateProduct(c *gin.Context, ctn *container.Container) {
	req := middlewares.GetValidatedModel(c).(*models.ProductCreateRequest)
	productModel := models.Product{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Quantity:    req.Quantity,
		CategoryID:  req.CategoryID,
		BrandID:     req.BrandID,
		Slug:        req.Slug,
		IsActive:    false,
	}
	if req.IsActive != nil {
		productModel.IsActive = *req.IsActive
	}
	newProduct, err := ctn.ProductService.CreateProduct(productModel)
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.SuccessResponse(c, http.StatusCreated, "Product created successfully", newProduct)
}

// UpdateProduct godoc
// @Summary Update product
// @Description Update product information (Admin only)
// @Tags products
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Product ID"
// @Param request body models.ProductUpdateRequest true "Product update data"
// @Success 200 {object} response.Response{data=models.SwaggerProduct} "Product updated successfully"
// @Failure 400 {object} response.Response "Invalid request"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 403 {object} response.Response "Forbidden"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /products/{id} [put]
func UpdateProduct(c *gin.Context, ctn *container.Container) {
	id := c.Param("id")
	req := middlewares.GetValidatedModel(c).(*models.ProductUpdateRequest)
	productModel := models.Product{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Quantity:    req.Quantity,
		CategoryID:  req.CategoryID,
		BrandID:     req.BrandID,
		Slug:        req.Slug,
		IsActive:    false,
	}
	if req.IsActive != nil {
		productModel.IsActive = *req.IsActive
	}
	updatedProduct, err := ctn.ProductService.UpdateProduct(id, productModel)
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.SuccessResponse(c, http.StatusOK, "Product updated successfully", updatedProduct)
}

// DeleteProduct godoc
// @Summary Delete product
// @Description Delete a product (Admin only)
// @Tags products
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Product ID"
// @Success 204 {object} response.Response "Product deleted successfully"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 403 {object} response.Response "Forbidden"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /products/{id} [delete]
func DeleteProduct(c *gin.Context, ctn *container.Container) {
	id := c.Param("id")
	err := ctn.ProductService.DeleteProduct(id)
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.SuccessResponse(c, http.StatusNoContent, "Product deleted successfully", nil)
}
