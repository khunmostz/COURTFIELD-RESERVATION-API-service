package authentication

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/khunmostz24/COURTFIELD-RESERVATION-API-service/model"
	"github.com/khunmostz24/COURTFIELD-RESERVATION-API-service/service"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func Register(c echo.Context, db *sql.DB) error {
	user := model.User{}

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "error", "data": "Bind data is error"})
	}
	user.Password = c.FormValue("password")

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "error", "data": "Failed to hash password"})
	}

	user.Password = string(hashedPassword)

	file, err := c.FormFile("image_url")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "error", "data": err})
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	cloudinaryInstance := service.Cloudinary{}

	cld, ctx := cloudinaryInstance.Credentials()
	user.Name = c.FormValue("name")
	user.Email = c.FormValue("email")
	user.Identification = c.FormValue("identification")
	user.ImageURL = cloudinaryInstance.UploadImage(cld, ctx, "profileUrl")

	rawQuery := `
	INSERT INTO users (image_url, username, email, password, identification)
	VALUES (?, ?, ?, ?, ?);
	`
	stmt, err := db.Prepare(rawQuery)
	if err != nil {
		log.Printf("Error preparing statement: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "prepare statement failed"})
	}
	r, err := stmt.Exec(user.ImageURL, user.Name, user.Email, user.Password, user.Identification)
	defer stmt.Close()
	switch {
	case err != nil:
		log.Printf("Error executing statement: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "execute statement failed"})
	case r == nil:
		log.Printf("No results returned: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "no result returned"})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{"message": "success", "data": user})
}

func Login(c echo.Context, db *sql.DB) error {
	user := model.User{}
	username := c.FormValue("username")
	password := c.FormValue("password")

	rawQuery := `
		SELECT * FROM users
		WHERE username=?;
	`
	rows, err := db.Query(rawQuery, username)

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

	if condition := checkPasswordHash(password, user.Password); condition == false {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Unauthorized"})

	}

	jwtInstance := service.JWTService{}
	jsonData, err := json.Marshal(user)
	if err != nil {
		log.Println(err)
	}
	token, err := jwtInstance.GenerateToken(string(jsonData))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to generate token"})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": map[string]interface{}{
			"user":  user,
			"token": token,
		},
	})
}
