package history

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
)


func GetAll(c echo.Context, db *sql.DB) error {
	return c.JSON(http.StatusOK, map[string]interface{}{"message": "success", "data": ""})
}