package users

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/khunmostz24/COURTFIELD-RESERVATION-API-service/model"
	"github.com/labstack/echo/v4"
)

func GetById(c echo.Context, db *sql.DB) error {
	user := model.User{}
	id := c.Param("id")
	if _, err := strconv.Atoi(id); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID format"})
	}
	rawQuery := `
	SELECT * FROM users WHERE user_id=?;
	`
	rows, err := db.Query(rawQuery, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
	}
	defer rows.Close()
	if rows.Next() {
		if err := rows.Scan(&user.ID, &user.ImageURL, &user.Name, &user.Email, &user.Password, &user.Identification, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
		}
	} else {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Not Found"})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"message": "success", "data": user})
}

func GetAll(c echo.Context, db *sql.DB) error {
	users := []model.User{}
	limit := c.QueryParam("limit")

	if limit == "" {
		rawQuery := `
		SELECT * FROM users;
		`
		rows, err := db.Query(rawQuery)
		if err != nil {
			log.Printf("Error query statement: %v", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
		}
		defer rows.Close()
		for rows.Next() {
			var user model.User
			if err := rows.Scan(&user.ID, &user.ImageURL, &user.Name, &user.Email, &user.Password, &user.Identification, &user.CreatedAt, &user.UpdatedAt); err != nil {
				log.Printf("Error scan: %v", err)
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
			}
			users = append(users, user)
		}
	} else {
		rawQuery := `
		SELECT * FROM users LIMIT ?
	`
		rows, err := db.Query(rawQuery, limit)
		if err != nil {
			log.Printf("Error query statement: %v", err)
			return c.JSON(http.StatusInternalServerError, err)
		}
		defer rows.Close()
		for rows.Next() {
			var user model.User
			if err := rows.Scan(&user.ID, &user.ImageURL, &user.Name, &user.Email, &user.Password, &user.Identification, &user.CreatedAt, &user.UpdatedAt); err != nil {
				log.Printf("Error scan: %v", err)
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
			}
			users = append(users, user)
		}

		if err := rows.Err(); err != nil {
			log.Printf("Error: %v", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"message": "success", "data": users})
}

func UpdateById(c echo.Context, db *sql.DB) error {
	return c.JSON(http.StatusOK, map[string]interface{}{"message": "success", "data": ""})
}
