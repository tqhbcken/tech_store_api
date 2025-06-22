package handlers

import (
	"api_techstore/internal/models"
	"api_techstore/internal/services"
	"api_techstore/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllCategories(c *gin.Context) {
	//goi toi service
	categories, err := services.GetAllCategories()
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.SuccessResponse(c, http.StatusOK, "Categories retrieved successfully", categories)
}

func GetCategoryById(c *gin.Context) {
	id := c.Param("id")

	//goi toi service
	category, err := services.GetCategoryById(id)
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.SuccessResponse(c, http.StatusOK, "Category retrieved successfully", category)
}

func CreateCategory(c *gin.Context) {
	//lay data tu body
	var category models.CategoryReq
	if err := c.ShouldBindJSON(&category); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, "Invalid input")
		return
	}

	// Chuyen CategoryReq sang Category
	categoryModel := models.Category{
		Name: category.Name,
		Slug: category.Slug,
	}

	//goi toi service
	newCategory, err := services.CreateCategory(categoryModel)
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.SuccessResponse(c, http.StatusCreated, "Category created successfully", newCategory)
}

func UpdateCategory(c *gin.Context) {
	id := c.Param("id")
	var category models.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, "Invalid input")
		return
	}

	//goi toi service
	updatedCategory, err := services.UpdateCategory(id, category)
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.SuccessResponse(c, http.StatusOK, "Category updated successfully", updatedCategory)
}

func DeleteCategory(c *gin.Context) {
	id := c.Param("id")

	//goi toi service
	err := services.DeleteCategory(id)
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.SuccessResponse(c, http.StatusNoContent, "Category deleted successfully", nil)
}
