package handlers

import (
	"api_techstore/internal/container"
	"api_techstore/internal/middlewares"
	"api_techstore/internal/models"
	"api_techstore/pkg/response"
	"net/http"
	"strconv"

	apperrors "api_techstore/pkg/errors"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CreateAddress godoc
// @Summary Create new address
// @Description Create a new address for the current user
// @Tags addresses
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body models.AddressCreateRequest true "Address data"
// @Success 201 {object} response.Response{data=models.SwaggerAddress} "Address created successfully"
// @Failure 400 {object} response.Response "Invalid request"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /addresses [post]
func CreateAddress(c *gin.Context, ctn *container.Container) {
	// Get user ID from context
	userID, exists := c.Get("user_id")
	if !exists {
		response.NewErrorResponse(c, apperrors.NewUnauthorized())
		return
	}

	// Convert user ID to uint
	userIDUint, ok := userID.(uint)
	if !ok {
		response.HandleError(c, apperrors.New(apperrors.ErrCodeInvalidInput, "Invalid user id type", http.StatusInternalServerError))
		return
	}

	// Get validated model from middleware
	req := middlewares.GetValidatedModel(c).(*models.AddressCreateRequest)

	address := models.Address{
		UserID:       userIDUint,
		FullName:     req.FullName,
		Phone:        req.Phone,
		AddressLine1: req.AddressLine1,
		AddressLine2: req.AddressLine2,
		City:         req.City,
		District:     req.District,
		IsDefault:    false,
	}

	if req.IsDefault != nil {
		address.IsDefault = *req.IsDefault
	}

	newAddress, err := ctn.AddressService.CreateAddress(address)
	if err != nil {
		response.DatabaseErrorResponse(c, err)
		return
	}

	response.SuccessResponse(c, http.StatusCreated, "Address created successfully", newAddress)
}

// GetAddresses godoc
// @Summary Get all addresses
// @Description Retrieve all addresses for the current user
// @Tags addresses
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} response.Response{data=[]models.SwaggerAddress} "Addresses retrieved successfully"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /addresses [get]
func GetAddresses(c *gin.Context, ctn *container.Container) {
	// Get user ID from context
	userID, exists := c.Get("user_id")
	if !exists {
		response.NewErrorResponse(c, apperrors.NewUnauthorized())
		return
	}

	// Convert user ID to uint
	userIDUint, ok := userID.(uint)
	if !ok {
		response.HandleError(c, apperrors.New(apperrors.ErrCodeInvalidInput, "Invalid user id type", http.StatusInternalServerError))
		return
	}

	addresses, err := ctn.AddressService.GetAllAddresses(userIDUint)
	if err != nil {
		response.DatabaseErrorResponse(c, err)
		return
	}

	response.SuccessResponse(c, http.StatusOK, "Addresses retrieved successfully", addresses)
}

// GetAddressByID godoc
// @Summary Get address by ID
// @Description Retrieve a specific address by ID
// @Tags addresses
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Address ID"
// @Success 200 {object} response.Response{data=models.SwaggerAddress} "Address retrieved successfully"
// @Failure 400 {object} response.Response "Invalid address id"
// @Failure 404 {object} response.Response "Address not found"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /addresses/{id} [get]
func GetAddressByID(c *gin.Context, ctn *container.Container) {
	id := c.Param("id")
	addressID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		response.NewErrorResponse(c, apperrors.NewValidationFailed("Invalid address id"))
		return
	}

	address, err := ctn.AddressService.GetAddressByID(uint(addressID))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			response.NotFoundResponse(c, "Address")
			return
		}
		response.DatabaseErrorResponse(c, err)
		return
	}

	response.SuccessResponse(c, http.StatusOK, "Address retrieved successfully", address)
}

// UpdateAddress godoc
// @Summary Update address
// @Description Update address information
// @Tags addresses
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Address ID"
// @Param request body models.AddressUpdateRequest true "Address update data"
// @Success 200 {object} response.Response{data=models.SwaggerAddress} "Address updated successfully"
// @Failure 400 {object} response.Response "Invalid request"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /addresses/{id} [put]
func UpdateAddress(c *gin.Context, ctn *container.Container) {
	id := c.Param("id")
	addressID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		response.NewErrorResponse(c, apperrors.NewValidationFailed("Invalid address id"))
		return
	}

	// Get validated model from middleware
	req := middlewares.GetValidatedModel(c).(*models.AddressUpdateRequest)

	address := models.Address{
		FullName:     req.FullName,
		Phone:        req.Phone,
		AddressLine1: req.AddressLine1,
		AddressLine2: req.AddressLine2,
		City:         req.City,
		District:     req.District,
		IsDefault:    false,
	}

	if req.IsDefault != nil {
		address.IsDefault = *req.IsDefault
	}

	updatedAddress, err := ctn.AddressService.UpdateAddress(uint(addressID), address)
	if err != nil {
		response.DatabaseErrorResponse(c, err)
		return
	}

	response.SuccessResponse(c, http.StatusOK, "Address updated successfully", updatedAddress)
}

// DeleteAddress godoc
// @Summary Delete address
// @Description Delete an address
// @Tags addresses
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Address ID"
// @Success 200 {object} response.Response "Address deleted successfully"
// @Failure 400 {object} response.Response "Invalid address id"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /addresses/{id} [delete]
func DeleteAddress(c *gin.Context, ctn *container.Container) {
	id := c.Param("id")
	addressID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		response.NewErrorResponse(c, apperrors.NewValidationFailed("Invalid address id"))
		return
	}

	err = ctn.AddressService.DeleteAddress(uint(addressID))
	if err != nil {
		response.DatabaseErrorResponse(c, err)
		return
	}

	response.SuccessResponse(c, http.StatusOK, "Address deleted successfully", nil)
}
