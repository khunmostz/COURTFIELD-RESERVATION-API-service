package booking

import (
	"net/http"

	"github.com/khunmostz24/COURTFIELD-RESERVATION-API-service/artifact/proxy/response"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Create(c echo.Context, db *gorm.DB) error {
	var booking Booking
	err := c.Bind(&booking)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.BaseResponse{
			Message: "error",
			Data:    err,
		})
	}

	if err := db.Create(&booking).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, response.BaseResponse{
			Message: "error",
			Data:    err,
		})
	}

	response := response.BaseResponse{
		Message: "success",
		Data:    booking,
	}

	return c.JSON(http.StatusOK, response)
}
