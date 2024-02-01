package sportequipment

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Create(c echo.Context, db *sql.DB) error {
	return c.JSON(http.StatusOK, map[string]interface{}{"message": "success", "data": ""})
}

func GetById(c echo.Context, db *sql.DB) error {
	return c.JSON(http.StatusOK, map[string]interface{}{"message": "success", "data": ""})
}

func GetAll(c echo.Context, db *sql.DB) error {
	return c.JSON(http.StatusOK, map[string]interface{}{"message": "success", "data": ""})
}

func UpdateById(c echo.Context, db *sql.DB) error {
	return c.JSON(http.StatusOK, map[string]interface{}{"message": "success", "data": ""})
}
