package handlers

import (
	"api_techstore/internal/container"
	"api_techstore/internal/middlewares"
	"api_techstore/internal/models"
	"api_techstore/pkg/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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
	req := middlewares.GetValidatedModel(c).(*models.AddressCreateRequest)
	// Lấy userID từ context (giả định đã có middleware JWT)
	userID, exists := c.Get("user_id")
	if !exists {
		response.ErrorResponse(c, http.StatusUnauthorized, "Unauthorized")
		return
	}
	uid, ok := userID.(uint)
	if !ok {
		response.ErrorResponse(c, http.StatusInternalServerError, "Invalid user id type")
		return
	}
	addressModel := models.Address{
		UserID:       uid, // luôn lấy từ context, không lấy từ req.UserID
		FullName:     req.FullName,
		Phone:        req.Phone,
		AddressLine1: req.AddressLine1,
		AddressLine2: req.AddressLine2,
		City:         req.City,
		District:     req.District,
		IsDefault:    false,
	}
	if req.IsDefault != nil {
		addressModel.IsDefault = *req.IsDefault
	}
	address, err := ctn.AddressService.CreateAddress(addressModel)
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.SuccessResponse(c, http.StatusCreated, "Address created successfully", address)
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
	userID, exists := c.Get("user_id")
	if !exists {
		response.ErrorResponse(c, http.StatusUnauthorized, "Unauthorized")
		return
	}
	uid, ok := userID.(uint)
	if !ok {
		response.ErrorResponse(c, http.StatusInternalServerError, "Invalid user id type")
		return
	}
	addresses, err := ctn.AddressService.GetAllAddresses(uid)
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
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
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, "Invalid address id")
		return
	}
	address, err := ctn.AddressService.GetAddressByID(uint(id))
	if err != nil {
		response.ErrorResponse(c, http.StatusNotFound, "Address not found")
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
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, "Invalid address id")
		return
	}
	req := middlewares.GetValidatedModel(c).(*models.AddressUpdateRequest)
	addressModel := models.Address{
		FullName:     req.FullName,
		Phone:        req.Phone,
		AddressLine1: req.AddressLine1,
		AddressLine2: req.AddressLine2,
		City:         req.City,
		District:     req.District,
		IsDefault:    false,
	}
	if req.IsDefault != nil {
		addressModel.IsDefault = *req.IsDefault
	}
	address, err := ctn.AddressService.UpdateAddress(uint(id), addressModel)
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.SuccessResponse(c, http.StatusOK, "Address updated successfully", address)
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
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, "Invalid address id")
		return
	}
	if err := ctn.AddressService.DeleteAddress(uint(id)); err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.SuccessResponse(c, http.StatusOK, "Address deleted successfully", nil)
}
