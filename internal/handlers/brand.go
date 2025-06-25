package handlers

import (
	"api_techstore/internal/container"
	"api_techstore/internal/middlewares"
	"api_techstore/internal/models"
	"api_techstore/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllBrands(c *gin.Context, ctn *container.Container) {
	brands, err := ctn.BrandService.GetAllBrands()
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.SuccessResponse(c, http.StatusOK, "Brands retrieved successfully", brands)
}

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
