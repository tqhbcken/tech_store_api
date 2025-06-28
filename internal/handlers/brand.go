package handlers

import (
	"api_techstore/internal/container"
	"api_techstore/internal/middlewares"
	"api_techstore/internal/models"
	"api_techstore/pkg/response"
	"net/http"

	apperrors "api_techstore/pkg/errors"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
		response.DatabaseErrorResponse(c, err)
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
// @Failure 400 {object} response.Response "Invalid request"
// @Failure 404 {object} response.Response "Brand not found"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /brands/{id} [get]
func GetBrandById(c *gin.Context, ctn *container.Container) {
	id := c.Param("id")
	if id == "" {
		response.NewErrorResponse(c, apperrors.NewValidationFailed("Brand ID is required"))
		return
	}

	brand, err := ctn.BrandService.GetBrandById(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			response.NotFoundResponse(c, "Brand")
			return
		}
		response.DatabaseErrorResponse(c, err)
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
	req := middlewares.GetValidatedModel(c).(*models.BrandCreateRequest)

	brand := models.Brand{
		Name:        req.Name,
		Description: req.Description,
		Slug:        req.Slug,
		IsActive:    false,
	}

	if req.IsActive != nil {
		brand.IsActive = *req.IsActive
	}

	newBrand, err := ctn.BrandService.CreateBrand(brand)
	if err != nil {
		response.DatabaseErrorResponse(c, err)
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
		response.NewErrorResponse(c, apperrors.NewValidationFailed("Brand ID is required"))
		return
	}

	// Check if brand exists
	brand, err := ctn.BrandService.GetBrandById(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			response.NotFoundResponse(c, "Brand")
			return
		}
		response.DatabaseErrorResponse(c, err)
		return
	}

	req := middlewares.GetValidatedModel(c).(*models.BrandUpdateRequest)

	updateBrand := models.Brand{
		Name:        req.Name,
		Description: req.Description,
		Slug:        req.Slug,
		IsActive:    brand.IsActive, // Keep existing value if not provided
	}

	if req.IsActive != nil {
		updateBrand.IsActive = *req.IsActive
	}

	updatedBrand, err := ctn.BrandService.UpdateBrand(id, updateBrand)
	if err != nil {
		response.DatabaseErrorResponse(c, err)
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
		response.NewErrorResponse(c, apperrors.NewValidationFailed("Brand ID is required"))
		return
	}

	// Check if brand exists
	_, err := ctn.BrandService.GetBrandById(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			response.NotFoundResponse(c, "Brand")
			return
		}
		response.DatabaseErrorResponse(c, err)
		return
	}

	err = ctn.BrandService.DeleteBrand(id)
	if err != nil {
		response.DatabaseErrorResponse(c, err)
		return
	}

	response.SuccessResponse(c, http.StatusOK, "Brand deleted successfully", nil)
}
