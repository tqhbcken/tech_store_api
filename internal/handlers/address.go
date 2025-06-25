package handlers

import (
	"api_techstore/internal/container"
	"api_techstore/internal/models"
	"api_techstore/pkg/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateAddress(c *gin.Context, ctn *container.Container) {
	var req models.Address
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, "Invalid input")
		return
	}
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
	req.UserID = uid
	address, err := ctn.AddressService.CreateAddress(req)
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.SuccessResponse(c, http.StatusCreated, "Address created successfully", address)
}

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

func UpdateAddress(c *gin.Context, ctn *container.Container) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, "Invalid address id")
		return
	}
	var req models.Address
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, "Invalid input")
		return
	}
	address, err := ctn.AddressService.UpdateAddress(uint(id), req)
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.SuccessResponse(c, http.StatusOK, "Address updated successfully", address)
}

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
