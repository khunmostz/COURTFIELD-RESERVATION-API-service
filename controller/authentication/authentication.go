package authentication

import (
	"net/http"

	"github.com/khunmostz24/COURTFIELD-RESERVATION-API-service/artifact/proxy/response"
	"github.com/khunmostz24/COURTFIELD-RESERVATION-API-service/service"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func Register(c echo.Context, db *gorm.DB) error {

	cloudinaryInstance := service.Cloudinary{}
	cld, ctx := cloudinaryInstance.Credentials()
	file, err := c.FormFile("imageUrl")
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Message: "error",
			Data:    err,
		})
	}

	hashedPassword, err := hashPassword(c.FormValue("password"))

	user := User{
		ImageURL:       cloudinaryInstance.UploadImage(cld, ctx, "productUrl", file),
		Name:           c.FormValue("username"),
		Password:       hashedPassword,
		Identification: c.FormValue("identification"),
	}

	// Check if the username already exists
	if err := db.Where("username = ?", c.FormValue("username")).First(&user).Error; err == nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Username already exists"})
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to hash password"})
	}

	if err := db.Create(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create user"})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "User registered successfully",
	})
}
