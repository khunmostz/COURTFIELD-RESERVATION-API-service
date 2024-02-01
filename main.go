package main

import (
	"log"
	"os"

	"github.com/khunmostz24/COURTFIELD-RESERVATION-API-service/controller/authentication"
	"github.com/khunmostz24/COURTFIELD-RESERVATION-API-service/controller/payment"
	"github.com/khunmostz24/COURTFIELD-RESERVATION-API-service/controller/sportequipment"

	_ "github.com/go-sql-driver/mysql"
	"github.com/khunmostz24/COURTFIELD-RESERVATION-API-service/controller/courtfield"
	"github.com/khunmostz24/COURTFIELD-RESERVATION-API-service/controller/history"
	"github.com/khunmostz24/COURTFIELD-RESERVATION-API-service/controller/role"
	"github.com/khunmostz24/COURTFIELD-RESERVATION-API-service/controller/users"
	"github.com/khunmostz24/COURTFIELD-RESERVATION-API-service/initialize"

	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func init() {
	initialize.Conn()
	initialize.LoadSQLFile("migration/CourtfieldReservation.sql")
}

func main() {
	e := echo.New()
	e.Use(middleware.CORS())
	port := os.Getenv("PORT")

	log.Println("SQL file loaded and executed.")

	e.Use(middleware.Logger())

	// e.GET("/", func(c echo.Context) error {
	// 	return c.JSON(http.StatusOK, map[string]interface{}{"message": "success", "data": "Hello world"})
	// })
	// Authentication
	// /auth/login-by-password POST (ล็อคอิน)
	e.POST("/authentication/login", func(c echo.Context) error {
		return authentication.Login(c, initialize.Db)
	})
	// /auth/register-by-password POST (สมัคร)
	e.POST("/authentication/register", func(c echo.Context) error {
		return authentication.Register(c, initialize.Db)
	})

	// Users
	// e.POST("/create/user/profile", func(c echo.Context) error {
	// 	return users.Create(c, initialize.Db)
	// })
	// /user/profile GET
	e.GET("/user/profile", func(c echo.Context) error {
		return users.GetAll(c, initialize.Db)
	})
	// /user/profile/:id GET (ดูโปรไฟล์)
	e.GET("/user/profile/:id", func(c echo.Context) error {
		return users.GetById(c, initialize.Db)
	})
	// /user/edit/profile/:id PUT (แก้ไขโปรไฟล์)
	e.GET("/user/edit/profile/:id", func(c echo.Context) error {
		return users.UpdateById(c, initialize.Db)
	})
	// Role
	// role/create/manager POST (สร้าง manager, admin)
	e.GET("/role/create/manager", func(c echo.Context) error {
		return role.Create(c, initialize.Db)
	})
	// role/edit/manager PUT (แก้ไข manager)
	e.GET("role/edit/manager", func(c echo.Context) error {
		return role.UpdateById(c, initialize.Db)
	})

	// Reservation and Booking
	// /reservation/calendar (ดูตารางการจอง)
	// /reservation/courtfield POST (จองสนาม)
	// /reservation/sport-equipment POST (จองอุปกรณ์)
	// /canceling/reservation?booking_id

	// History
	// /history?type=history GET (ดูประวัติ)
	e.GET("/history?type=history", func(c echo.Context) error {
		return history.GetAll(c, initialize.Db)
	})

	// Payments
	// /payments/prompt-pay POST (จ่ายเงิน)
	e.POST("/payments/prompt-pay", func(c echo.Context) error {
		return payment.PaymentQrCode(c, initialize.Db)
	})
	// /payments/prompt-pay?id=?
	e.GET("/payments/prompt-pay", func(c echo.Context) error {
		return payment.GetPaymentStatus(c, initialize.Db)
	})

	// Court
	// /get/courts GET (ดูสนามทั้งหมด)
	e.GET("/get/courts", func(c echo.Context) error {
		return courtfield.GetAll(c, initialize.Db)
	})
	// /get/courts?=id
	e.GET("/get/court/:id", func(c echo.Context) error {
		return courtfield.GetById(c, initialize.Db)
	})
	// /create/court POST (สร้างสนาม)
	e.POST("/create/court", func(c echo.Context) error {
		return courtfield.Create(c, initialize.Db)
	})
	// /edit/court?court_id=? (แก้ไขสนาม)
	e.PUT("/put/court", func(c echo.Context) error {
		return courtfield.UpdateById(c, initialize.Db)
	})

	// Sport-equipment
	// /create/sport-equipment (สร้างอุปกรณ์กีฬา)
	e.POST("/create/sport-equipment", func(c echo.Context) error {
		return sportequipment.Create(c, initialize.Db)
	})
	// /edit/sport-equipment (สร้างอุปกรณ์กีฬา)
	e.PUT("/put/sport-equipment", func(c echo.Context) error {
		return sportequipment.UpdateById(c, initialize.Db)
	})

	log.Printf("Starting server on port %s\n", port)
	log.Fatal(e.Start(":" + port))

}
