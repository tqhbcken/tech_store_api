package handlers

import (
	"api_techstore/internal/container"
	"api_techstore/internal/middlewares"
	"api_techstore/internal/models"
	"api_techstore/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAllBrands godoc
// @Summary Get all brands
// @Description Retrieve all brands
// @Tags brands
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} response.Response{data=[]models.SwaggerBrand} "Brands retrieved successfully"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /brands [get]
func GetAllBrands(c *gin.Context, ctn *container.Container) {
	brands, err := ctn.BrandService.GetAllBrands()
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.SuccessResponse(c, http.StatusOK, "Brands retrieved successfully", brands)
}

// GetBrandById godoc
// @Summary Get brand by ID
// @Description Retrieve a specific brand by ID
// @Tags brands
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Brand ID"
// @Success 200 {object} response.Response{data=models.SwaggerBrand} "Brand retrieved successfully"
// @Failure 404 {object} response.Response "Brand not found"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /brands/{id} [get]
func GetBrandById(c *gin.Context, ctn *container.Container) {
	id := c.Param("id")
	if id == "" {
		response.ErrorResponse(c, http.StatusBadRequest, "Brand ID is required")
		return
	}
	// Có thể kiểm tra id là số dương nếu cần
	brand, err := ctn.BrandService.GetBrandById(id)
	if err != nil || brand.ID == 0 {
		response.ErrorResponse(c, http.StatusNotFound, "Brand not found")
		return
	}
	response.SuccessResponse(c, http.StatusOK, "Brand retrieved successfully", brand)
}

// CreateBrand godoc
// @Summary Create new brand
// @Description Create a new brand (Admin only)
// @Tags brands
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body models.BrandCreateRequest true "Brand data"
// @Success 201 {object} response.Response{data=models.SwaggerBrand} "Brand created successfully"
// @Failure 400 {object} response.Response "Invalid request"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 403 {object} response.Response "Forbidden"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /brands [post]
func CreateBrand(c *gin.Context, ctn *container.Container) {
	// Lấy validated model từ middleware
	req := middlewares.GetValidatedModel(c).(*models.BrandCreateRequest)

	brandModel := models.Brand{
		Name:        req.Name,
		Description: req.Description,
		IsActive:    false,
		Slug:        req.Slug,
	}
	if req.IsActive != nil {
		brandModel.IsActive = *req.IsActive
	}

	newBrand, err := ctn.BrandService.CreateBrand(brandModel)
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.SuccessResponse(c, http.StatusCreated, "Brand created successfully", newBrand)
}

// UpdateBrand godoc
// @Summary Update brand
// @Description Update brand information (Admin only)
// @Tags brands
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Brand ID"
// @Param request body models.BrandUpdateRequest true "Brand update data"
// @Success 200 {object} response.Response{data=models.SwaggerBrand} "Brand updated successfully"
// @Failure 400 {object} response.Response "Invalid request"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 403 {object} response.Response "Forbidden"
// @Failure 404 {object} response.Response "Brand not found"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /brands/{id} [put]
func UpdateBrand(c *gin.Context, ctn *container.Container) {
	id := c.Param("id")
	if id == "" {
		response.ErrorResponse(c, http.StatusBadRequest, "Brand ID is required")
		return
	}

	_, err := ctn.BrandService.GetBrandById(id)
	if err != nil {
		response.ErrorResponse(c, http.StatusNotFound, "Brand not found")
		return
	}

	req := middlewares.GetValidatedModel(c).(*models.BrandUpdateRequest)
	brandModel := models.Brand{
		Name:        req.Name,
		Description: req.Description,
		IsActive:    false,
		Slug:        req.Slug,
	}
	if req.IsActive != nil {
		brandModel.IsActive = *req.IsActive
	}

	updatedBrand, err := ctn.BrandService.UpdateBrand(id, brandModel)
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.SuccessResponse(c, http.StatusOK, "Brand updated successfully", updatedBrand)
}

// DeleteBrand godoc
// @Summary Delete brand
// @Description Delete a brand (Admin only)
// @Tags brands
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Brand ID"
// @Success 200 {object} response.Response "Brand deleted successfully"
// @Failure 400 {object} response.Response "Invalid request"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 403 {object} response.Response "Forbidden"
// @Failure 404 {object} response.Response "Brand not found"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /brands/{id} [delete]
func DeleteBrand(c *gin.Context, ctn *container.Container) {
	id := c.Param("id")

	if id == "" {
		response.ErrorResponse(c, http.StatusBadRequest, "Brand ID is required")
		return
	}

	_, err := ctn.BrandService.GetBrandById(id)
	if err != nil {
		response.ErrorResponse(c, http.StatusNotFound, "Brand not found")
		return
	}

	err = ctn.BrandService.DeleteBrand(id)
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.SuccessResponse(c, http.StatusOK, "Brand deleted successfully", nil)
}
