package authentication

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
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
		Name:           c.FormValue("username"),
		Password:       hashedPassword,
		Email:          c.FormValue("email"),
		Identification: c.FormValue("identification"),
	}

	// Check if the username already exists
	if err := db.Where("name = ?", c.FormValue("username")).First(&user).Error; err == nil {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Message: "error",
			Data:    "Username already exists",
		})
	}

	user.ImageURL = cloudinaryInstance.UploadImage(cld, ctx, "profileUrl", file)

	if err := db.Create(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, response.BaseResponse{
			Message: "error",
			Data:    err,
		})
	}

	response := response.BaseResponse{
		Message: "success",
		Data:    user,
	}

	return c.JSON(http.StatusOK, response)
}

func Login(c echo.Context, db *gorm.DB) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	// Find the user by username
	var user User
	if err := db.Where("name = ?", username).First(&user).Error; err != nil {
		return c.JSON(http.StatusUnauthorized, response.BaseResponse{
			Message: "error",
			Data:    "Invalid username",
		})
	}

	// Check the password
	if !checkPasswordHash(password, user.Password) {
		return c.JSON(http.StatusUnauthorized, response.BaseResponse{
			Message: "error",
			Data:    "Invalid password",
		})
	}

	jwtInstance := service.JWTService{}

	claims := &jwtCustomClaims{
		ID:             user.ID,
		ImageURL:       user.ImageURL,
		Name:           user.Name,
		Email:          user.Email,
		Identification: user.Identification,
		CreatedAt:      user.CreatedAt,
		UpdatedAt:      user.UpdatedAt,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(), // 3 days
		},
	}

	// Generate JWT token
	token, err := jwtInstance.GenerateToken(claims)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.BaseResponse{
			Message: "error",
			Data:    err,
		})
	}

	response := response.BaseResponse{
		Message: "success",
		Data: Credentials{
			User:  user,
			Token: token,
		},
	}

	return c.JSON(http.StatusOK, response)
}
