package courtfield

import (
	"database/sql"
	"net/http"
	"strconv"
	"strings"

	"github.com/khunmostz24/COURTFIELD-RESERVATION-API-service/model"
	"github.com/labstack/echo/v4"
)

func GetById(c echo.Context, db *sql.DB) error {
	courts := model.Court{}
	id := c.Param("id")
	if _, err := strconv.Atoi(id); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID format"})
	}
	rawQuery := `
	SELECT * FROM courts WHERE courts_id=?;
	`
	rows, err := db.Query(rawQuery, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
	}
	defer rows.Close()
	if rows.Next() {
		if err := rows.Scan(&courts.ID, &courts.Name, &courts.Description, &courts.Price, &courts.CreatedAt, &courts.UpdatedAt); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
		}
	} else {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Not Found"})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"message": "success", "data": courts})
}

func GetAll(c echo.Context, db *sql.DB) error {
	courts := []model.Court{}
	limit := c.QueryParam("limit")

	if limit == "" {
		rawQuery := `
		SELECT * FROM courts;
		`
		rows, err := db.Query(rawQuery)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
		}
		defer rows.Close()
		for rows.Next() {
			var court model.Court
			if err := rows.Scan(&court.ID, &court.Name, &court.Description, &court.Price, &court.CreatedAt, &court.UpdatedAt); err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
			}
			courts = append(courts, court)
		}
	} else {
		rawQuery := `
		SELECT * FROM courts LIMIT ?
	`
		rows, err := db.Query(rawQuery, limit)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		defer rows.Close()
		for rows.Next() {
			var court model.Court
			if err := rows.Scan(&court.ID, &court.Name, &court.Description, &court.Price, &court.CreatedAt, &court.UpdatedAt); err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
			}
			courts = append(courts, court)
		}

		if err := rows.Err(); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"message": "success", "data": courts})
}

func Create(c echo.Context, db *sql.DB) error {
	court := model.Court{}

	if err := c.Bind(&court); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	rawQuery := `
	INSERT INTO courts (courts_name,courts_description,price)
	VALUES (?, ?, ?);
	`
	stmt, err := db.Prepare(rawQuery)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
	}

	r, err := stmt.Exec(court.Name, court.Description, court.Price)
	defer stmt.Close()
	switch {
	case err != nil:
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
	case err == nil:
		id, _ := r.LastInsertId()
		court.ID = int(id)
		return c.JSON(http.StatusCreated, court)
	default:
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
	}

}

func UpdateById(c echo.Context, db *sql.DB) error {
	courtIDStr := c.QueryParam("id")
	if courtIDStr == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "ID is required"})
	}

	// Convert courtIDStr to int
	courtID, err := strconv.Atoi(courtIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}

	// Retrieve existing court data from the database
	existingCourt := model.Court{}
	err = db.QueryRow("SELECT * FROM courts WHERE courts_id=?", courtID).
		Scan(&existingCourt.ID, &existingCourt.Name, &existingCourt.Description, &existingCourt.Price, &existingCourt.CreatedAt, &existingCourt.UpdatedAt)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
	}

	// Update only the fields provided in the request
	if err := c.Bind(&existingCourt); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	// Build the SQL query based on the provided fields
	updateQuery := "UPDATE courts SET "
	var params []interface{}

	if existingCourt.Name != "" {
		updateQuery += "courts_name=?,"
		params = append(params, existingCourt.Name)
	}

	if existingCourt.Description != "" {
		updateQuery += "courts_description=?,"
		params = append(params, existingCourt.Description)
	}

	if existingCourt.Price != 0 {
		updateQuery += "price=?,"
		params = append(params, existingCourt.Price)
	}

	// Remove the trailing comma from the query
	updateQuery = strings.TrimSuffix(updateQuery, ",")

	// Add the WHERE clause to specify the court ID
	updateQuery += " WHERE courts_id=?"

	// Add the court ID to the parameters
	params = append(params, courtID)

	// Prepare and execute the SQL statement
	stmt, err := db.Prepare(updateQuery)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
	}

	_, err = stmt.Exec(params...)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"message": "success", "data": existingCourt})
}
