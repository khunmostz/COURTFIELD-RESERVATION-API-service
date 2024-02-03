package main

import (
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/khunmostz24/COURTFIELD-RESERVATION-API-service/controller/authentication"
	"github.com/khunmostz24/COURTFIELD-RESERVATION-API-service/controller/booking"
	"github.com/khunmostz24/COURTFIELD-RESERVATION-API-service/controller/payment"
	"github.com/khunmostz24/COURTFIELD-RESERVATION-API-service/controller/products"
	"github.com/khunmostz24/COURTFIELD-RESERVATION-API-service/initialize"

	_ "github.com/joho/godotenv/autoload"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func init() {
	initialize.Conn()
}

func main() {
	e := echo.New()
	e.Use(middleware.CORS())
	port := os.Getenv("PORT")

	initialize.Db.AutoMigrate(&products.Product{}, &products.Category{}, &authentication.User{}, &booking.Booking{}, &booking.BookingProduct{}, &payment.Payment{})

	authMiddleware := echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	})

	r := e.Group("/api")
	r.Use(authMiddleware)

	e.POST("/register", func(c echo.Context) error {
		return authentication.Register(c, initialize.Db)
	})

	e.POST("/login", func(c echo.Context) error {
		return authentication.Login(c, initialize.Db)
	})

	// r.POST("/validate-token", func(c echo.Context) error {
	// 	return authentication.Login(c, initialize.Db)
	// })

	r.POST("/create/product", func(c echo.Context) error {
		return products.Create(c, initialize.Db)
	})

	r.GET("/products", func(c echo.Context) error {
		return products.GetAll(c, initialize.Db)
	})

	r.GET("/product/:id", func(c echo.Context) error {
		return products.GetById(c, initialize.Db)
	})

	r.PUT("/product/:id", func(c echo.Context) error {
		return products.UpdateById(c, initialize.Db)
	})

	r.DELETE("/product/:id", func(c echo.Context) error {
		return products.DeleteById(c, initialize.Db)
	})

	r.POST("/create/booking", func(c echo.Context) error {
		return booking.Create(c, initialize.Db)
	})

	e.POST("/create/payment", func(c echo.Context) error {
		return payment.PaymentQrCode(c, initialize.Db)
	})

	e.GET("/create/payment/:id", func(c echo.Context) error {
		return payment.GetPaymentStatus(c, initialize.Db)
	})

	log.Printf("Starting server on port %s\n", port)
	log.Fatal(e.Start(":" + port))

}
