package handlers

import (
	"api_techstore/internal/container"
	"api_techstore/internal/middlewares"
	"api_techstore/internal/models"
	"api_techstore/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAllCategories godoc
// @Summary Get all categories
// @Description Retrieve all categories
// @Tags categories
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} response.Response{data=[]models.SwaggerCategory} "Categories retrieved successfully"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /categories [get]
func GetAllCategories(c *gin.Context, ctn *container.Container) {
	categories, err := ctn.CategoryService.GetAllCategories()
	if err != nil {
		response.DatabaseErrorResponse(c, err)
		return
	}
	response.SuccessResponse(c, http.StatusOK, "Categories retrieved successfully", categories)
}

// GetCategoryById godoc
// @Summary Get category by ID
// @Description Retrieve a specific category by ID
// @Tags categories
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Category ID"
// @Success 200 {object} response.Response{data=models.SwaggerCategory} "Category retrieved successfully"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /categories/{id} [get]
func GetCategoryById(c *gin.Context, ctn *container.Container) {
	id := c.Param("id")
	category, err := ctn.CategoryService.GetCategoryById(id)
	if err != nil {
		response.DatabaseErrorResponse(c, err)
		return
	}
	response.SuccessResponse(c, http.StatusOK, "Category retrieved successfully", category)
}

// CreateCategory godoc
// @Summary Create new category
// @Description Create a new category (Admin only)
// @Tags categories
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body models.CategoryCreateRequest true "Category data"
// @Success 201 {object} response.Response{data=models.SwaggerCategory} "Category created successfully"
// @Failure 400 {object} response.Response "Invalid request"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 403 {object} response.Response "Forbidden"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /categories [post]
func CreateCategory(c *gin.Context, ctn *container.Container) {
	req := middlewares.GetValidatedModel(c).(*models.CategoryCreateRequest)

	category := models.Category{
		Name: req.Name,
		Slug: req.Slug,
	}

	newCategory, err := ctn.CategoryService.CreateCategory(category)
	if err != nil {
		response.DatabaseErrorResponse(c, err)
		return
	}

	response.SuccessResponse(c, http.StatusCreated, "Category created successfully", newCategory)
}

// UpdateCategory godoc
// @Summary Update category
// @Description Update category information (Admin only)
// @Tags categories
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Category ID"
// @Param request body models.CategoryUpdateRequest true "Category update data"
// @Success 200 {object} response.Response{data=models.SwaggerCategory} "Category updated successfully"
// @Failure 400 {object} response.Response "Invalid request"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 403 {object} response.Response "Forbidden"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /categories/{id} [put]
func UpdateCategory(c *gin.Context, ctn *container.Container) {
	id := c.Param("id")
	req := middlewares.GetValidatedModel(c).(*models.CategoryUpdateRequest)

	category := models.Category{
		Name: req.Name,
		Slug: req.Slug,
	}

	updatedCategory, err := ctn.CategoryService.UpdateCategory(id, category)
	if err != nil {
		response.DatabaseErrorResponse(c, err)
		return
	}

	response.SuccessResponse(c, http.StatusOK, "Category updated successfully", updatedCategory)
}

// DeleteCategory godoc
// @Summary Delete category
// @Description Delete a category (Admin only)
// @Tags categories
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Category ID"
// @Success 204 {object} response.Response "Category deleted successfully"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 403 {object} response.Response "Forbidden"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /categories/{id} [delete]
func DeleteCategory(c *gin.Context, ctn *container.Container) {
	id := c.Param("id")
	err := ctn.CategoryService.DeleteCategory(id)
	if err != nil {
		response.DatabaseErrorResponse(c, err)
		return
	}

	response.SuccessResponse(c, http.StatusNoContent, "Category deleted successfully", nil)
}
