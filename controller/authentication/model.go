package authentication

import (
	"time"
)

type User struct {
	ID             int       `json:"user_id" gorm:"primaryKey;autoIncrement"`
	ImageURL       string    `json:"image_url,omitempty" gorm:"type:varchar(255)"`
	Name           string    `json:"name" gorm:"type:varchar(255)"`
	Email          string    `json:"email" gorm:"type:varchar(255)" `
	Password       string    `json:"-" gorm:"type:varchar(255)"` // for disable show on return model
	Identification string    `json:"identification" gorm:"type:varchar(255)"`
	CreatedAt      time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt      time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
