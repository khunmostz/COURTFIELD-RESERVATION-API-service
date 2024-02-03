package booking

import (
	"time"
)

type Booking struct {
	BookingID       int              `gorm:"primaryKey;autoIncrement" json:"bookingId"`
	UserID          int              `gorm:"int;not null" json:"userId"`
	BookingProducts []BookingProduct `gorm:"foreignKey:BookingID" json:"bookingProducts"`
	StartTime       time.Time        `gorm:"not null" json:"startTime"`
	EndTime         time.Time        `gorm:"not null" json:"endTime"`
	Price           float64          `gorm:"type:decimal(10,2)" json:"price"`
	PaymentStatus   string           `gorm:"varchar(100)" json:"paymentStatus"`
	CreatedAt       time.Time        `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt       time.Time        `gorm:"autoCreateTime" json:"updatedAt"`
}

type BookingProduct struct {
	ID        int       `json:"id" gorm:"primaryKey;autoIncrement"`
	BookingID int       `json:"bookingId" gorm:"type:int(11);not null"`
	ProductID int       `json:"productId" gorm:"type:int(11);not null"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoCreateTime"`
}
