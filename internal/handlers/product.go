package handlers

import (
	"api_techstore/internal/container"
	"api_techstore/internal/middlewares"
	"api_techstore/internal/models"
	"api_techstore/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllProducts(c *gin.Context, ctn *container.Container) {
	products, err := ctn.ProductService.GetAllProducts()
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.SuccessResponse(c, http.StatusOK, "Products retrieved successfully", products)
}

func GetProductById(c *gin.Context, ctn *container.Container) {
	id := c.Param("id")
	product, err := ctn.ProductService.GetProductById(id)
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.SuccessResponse(c, http.StatusOK, "Product retrieved successfully", product)
}

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

func DeleteProduct(c *gin.Context, ctn *container.Container) {
	id := c.Param("id")
	err := ctn.ProductService.DeleteProduct(id)
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.SuccessResponse(c, http.StatusNoContent, "Product deleted successfully", nil)
}
