package handlers

import (
	"api_techstore/internal/container"
	"api_techstore/internal/middlewares"
	"api_techstore/internal/models"
	"api_techstore/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllCategories(c *gin.Context, ctn *container.Container) {
	categories, err := ctn.CategoryService.GetAllCategories()
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.SuccessResponse(c, http.StatusOK, "Categories retrieved successfully", categories)
}

func GetCategoryById(c *gin.Context, ctn *container.Container) {
	id := c.Param("id")
	category, err := ctn.CategoryService.GetCategoryById(id)
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.SuccessResponse(c, http.StatusOK, "Category retrieved successfully", category)
}

func CreateCategory(c *gin.Context, ctn *container.Container) {
	req := middlewares.GetValidatedModel(c).(*models.CategoryCreateRequest)
	categoryModel := models.Category{
		Name: req.Name,
		Slug: req.Slug,
	}
	newCategory, err := ctn.CategoryService.CreateCategory(categoryModel)
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.SuccessResponse(c, http.StatusCreated, "Category created successfully", newCategory)
}

func UpdateCategory(c *gin.Context, ctn *container.Container) {
	id := c.Param("id")
	req := middlewares.GetValidatedModel(c).(*models.CategoryUpdateRequest)
	categoryModel := models.Category{
		Name: req.Name,
		Slug: req.Slug,
	}
	updatedCategory, err := ctn.CategoryService.UpdateCategory(id, categoryModel)
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.SuccessResponse(c, http.StatusOK, "Category updated successfully", updatedCategory)
}

func DeleteCategory(c *gin.Context, ctn *container.Container) {
	id := c.Param("id")
	err := ctn.CategoryService.DeleteCategory(id)
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.SuccessResponse(c, http.StatusNoContent, "Category deleted successfully", nil)
}
