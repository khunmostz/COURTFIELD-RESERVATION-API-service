package products

import (
	"time"
)

type Product struct {
	ProductID          int        `json:"id" gorm:"primaryKey;autoIncrement"`
	ImageURL           string     `json:"imageUrl" gorm:"type:varchar(255)"`
	ProductName        string     `json:"productName" gorm:"type:varchar(255)"`
	ProductDescription string     `json:"productDescription" gorm:"type:varchar(255)"`
	Price              int        `json:"price" gorm:"type:int(11)"`
	Category           []Category `json:"category" gorm:"foreignKey:ProductID"`
	CreatedAt          time.Time  `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt          time.Time  `json:"updatedAt" gorm:"autoUpdateTime"`
}

type Category struct {
	CategoryID   int       `json:"categoryId" gorm:"primaryKey;autoIncrement"`
	ProductID    int       `json:"productId" gorm:"type:int(11);not null"`
	CategoryName string    `json:"categoryName" gorm:"type:varchar(255);not null"`
	CreatedAt    time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt    time.Time `json:"updatedAt" gorm:"autoCreateTime"`
}
