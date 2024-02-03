package products

import (
	"net/http"
	"strconv"

	"github.com/khunmostz24/COURTFIELD-RESERVATION-API-service/artifact/proxy/response"
	"github.com/khunmostz24/COURTFIELD-RESERVATION-API-service/service"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Create(c echo.Context, db *gorm.DB) error {
	file, err := c.FormFile("imageUrl")
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Message: "error",
			Data:    err,
		})
	}
	cloudinaryInstance := service.Cloudinary{}
	cld, ctx := cloudinaryInstance.Credentials()
	imageUrl := cloudinaryInstance.UploadImage(cld, ctx, "productUrl", file)

	println("imageUrl: %s", imageUrl)

	parsePrice, err := strconv.Atoi(c.FormValue("price"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, &response.BaseResponse{
			Message: "error",
			Data:    err.Error(),
		})
	}

	category := Category{
		CategoryName: c.FormValue("categoryName"),
	}

	product := Product{
		ImageURL:           imageUrl,
		ProductName:        c.FormValue("productName"),
		ProductDescription: c.FormValue("productDescription"),
		Price:              parsePrice,
		Category: []Category{
			category,
		},
	}

	if err := db.Create(&product).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, response.BaseResponse{
			Message: "error",
			Data:    err,
		})
	}

	response := response.BaseResponse{
		Message: "success",
		Data:    product,
	}

	return c.JSON(http.StatusOK, response)
}
func GetAll(c echo.Context, db *gorm.DB) error {
	var products []Product

	if err := db.Preload("Category").Find(&products).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, response.BaseResponse{
			Message: "error",
			Data:    err,
		})
	}

	response := response.BaseResponse{
		Message: "success",
		Data:    products,
	}

	return c.JSON(http.StatusOK, response)
}

func GetById(c echo.Context, db *gorm.DB) error {
	id := c.Param("id")

	var product Product
	if err := db.Preload("Category").Find(&product, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, response.BaseResponse{
				Message: "error",
				Data:    "Product not found",
			})
		}
		return c.JSON(http.StatusInternalServerError, response.BaseResponse{
			Message: "error",
			Data:    err,
		})
	}

	response := response.BaseResponse{
		Message: "success",
		Data:    product,
	}

	return c.JSON(http.StatusOK, response)
}

func UpdateById(c echo.Context, db *gorm.DB) error {
	id := c.Param("id")

	// Parse the ID to an integer
	productID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Message: "error",
			Data:    "Invalid product ID",
		})
	}

	// Get the existing product from the database
	var existingProduct Product
	if err := db.Preload("Category").First(&existingProduct, productID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, response.BaseResponse{
				Message: "error",
				Data:    "Product not found",
			})
		}
		return c.JSON(http.StatusInternalServerError, response.BaseResponse{
			Message: "error",
			Data:    err,
		})
	}

	// Update the product with the new data
	// You can access the request body using c.Bind() or c.BindJSON() methods
	// Update the fields of existingProduct with the new data
	if err := c.Bind(&existingProduct); err != nil {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Message: "error",
			Data:    "Invalid request body",
		})
	}

	// Save the updated product to the database
	if err := db.Save(&existingProduct).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, response.BaseResponse{
			Message: "error",
			Data:    err,
		})
	}

	response := response.BaseResponse{
		Message: "success",
		Data:    existingProduct,
	}

	return c.JSON(http.StatusOK, response)
}

func DeleteById(c echo.Context, db *gorm.DB) error {
	// Get the product ID from the request parameters
	productID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.BaseResponse{
			Message: "error",
			Data:    err,
		})
	}

	// Delete the product with the given ID from the database
	err = db.Unscoped().Delete(&Product{}, productID).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.BaseResponse{
			Message: "error",
			Data:    err,
		})
	}

	// Return a success response
	return c.JSON(http.StatusOK, response.BaseResponse{
		Message: "success",
		Data:    "Product deleted successfully",
	})
}
