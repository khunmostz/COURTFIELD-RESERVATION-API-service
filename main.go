package main

import (
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/khunmostz24/COURTFIELD-RESERVATION-API-service/controller/authentication"
	"github.com/khunmostz24/COURTFIELD-RESERVATION-API-service/controller/products"
	"github.com/khunmostz24/COURTFIELD-RESERVATION-API-service/initialize"

	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func init() {
	initialize.Conn()
	// initialize.LoadSQLFile("migration/db.sql")
}

func main() {
	e := echo.New()
	e.Use(middleware.CORS())
	port := os.Getenv("PORT")

	initialize.Db.AutoMigrate(&products.Product{}, &products.Category{}, &authentication.User{})

	// authMiddleware := middleware.JWTWithConfig(middleware.JWTConfig{
	// 	SigningKey: []byte(os.Getenv("JWT_SECRET")),
	// })

	e.POST("/create/product", func(c echo.Context) error {
		return products.Create(c, initialize.Db)
	})

	e.GET("/products", func(c echo.Context) error {
		return products.GetAll(c, initialize.Db)
	})

	e.GET("/product/:id", func(c echo.Context) error {
		return products.GetById(c, initialize.Db)
	})

	e.PUT("/product/:id", func(c echo.Context) error {
		return products.UpdateById(c, initialize.Db)
	})

	e.DELETE("/product/:id", func(c echo.Context) error {
		return products.DeleteById(c, initialize.Db)
	})

	log.Printf("Starting server on port %s\n", port)
	log.Fatal(e.Start(":" + port))

}
