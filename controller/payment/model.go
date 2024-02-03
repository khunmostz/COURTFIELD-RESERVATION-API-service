package payment

import (
	"time"

	"github.com/khunmostz24/COURTFIELD-RESERVATION-API-service/controller/booking"
)

type Payment struct {
	PaymentID       string                   `json:"paymentId" gorm:"type:varchar(255);primaryKey;uniqueIndex:idx_payment_id"`
	BookingProducts []booking.BookingProduct `gorm:"foreignKey:PaymentID" json:"bookingProducts"`
	Source          string                   `json:"source" gorm:"type:varchar(255)"`
	Amount          int                      `json:"amount" gorm:"type:int(11)"`
	PaymentMethod   string                   `json:"paymentMethod" gorm:"type:varchar(255)"`
	Status          string                   `json:"status" gorm:"type:varchar(255)"`
	CreatedAt       time.Time                `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt       time.Time                `json:"updatedAt" gorm:"autoUpdateTime"`
}
