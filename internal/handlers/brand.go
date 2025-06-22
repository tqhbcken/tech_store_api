package handlers

import (
	"api_techstore/internal/models"
	"api_techstore/internal/services"
	"api_techstore/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllBrands(c *gin.Context) {
	//goi toi service
	brands, err := services.GetAllBrands()
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.SuccessResponse(c, http.StatusOK, "Brands retrieved successfully", brands)
}

func GetBrandById(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		response.ErrorResponse(c, http.StatusBadRequest, "Brand ID is required")
		return
	}

	idChecker, err := services.GetBrandById(id)
	if idChecker.ID == 0 || err != nil {
		response.ErrorResponse(c, http.StatusNotFound, "Brand not found")
		return
	}

	//goi toi service
	brand, err := services.GetBrandById(id)
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.SuccessResponse(c, http.StatusOK, "Brand retrieved successfully", brand)
}

func CreateBrand(c *gin.Context) {
	//lay data tu body
	var brand models.CreateBrandReq
	// Kiem tra du lieu nhap vao
	if err := c.ShouldBindJSON(&brand); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, "Invalid input: "+err.Error())
		return
	}

	//chuyen doi brandReq sang brand
	brandModel := models.Brand{
		Name:        brand.Name,
		Description: brand.Description,
		IsActive:    brand.IsActive,	
		Slug:        brand.Slug,
	}

	//goi toi service
	newBrand, err := services.CreateBrand(brandModel)
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.SuccessResponse(c, http.StatusCreated, "Brand created successfully", newBrand)
}

func UpdateBrand(c *gin.Context) {
	id := c.Param("id")
	var brand models.UpdateBrandReq
	//lay data tu body
	if err := c.ShouldBindJSON(&brand); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, "Invalid input: "+err.Error())
		return
	}

	//kiem tra id
	if id == "" {
		response.ErrorResponse(c, http.StatusBadRequest, "Brand ID is required")
		return
	}

	idChecker, err := services.GetBrandById(id)
	if idChecker.ID == 0 || err != nil {
		response.ErrorResponse(c, http.StatusNotFound, "Brand not found")
		return
	}

	//chuyen doi BrandReq sang Brand
	brandModel := models.Brand{
		Name:        brand.Name,	
		Description: brand.Description,
		IsActive:    brand.IsActive,
		Slug:        brand.Slug,
	}

	//goi toi service
	updatedBrand, err := services.UpdateBrand(id, brandModel)
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.SuccessResponse(c, http.StatusOK, "Brand updated successfully", updatedBrand)
}

func DeleteBrand(c *gin.Context) {
	id := c.Param("id")

	//kiem tra id
	if id == "" {
		response.ErrorResponse(c, http.StatusBadRequest, "Brand ID is required")
		return
	}

	idChecker, error := services.GetBrandById(id)
	if idChecker.ID == 0 || error != nil {
		response.ErrorResponse(c, http.StatusNotFound, "Brand not found")
		return
	}

	//goi toi service
	err := services.DeleteBrand(id)
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.SuccessResponse(c, http.StatusOK, "Brand deleted successfully", nil)
}

