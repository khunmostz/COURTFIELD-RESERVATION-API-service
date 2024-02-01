package model

import (
	"time"
)

type User struct {
	ID             int       `json:"user_id"`
	ImageURL       string    `json:"image_url,omitempty"`
	Name           string    `json:"name"`
	Email          string    `json:"email"`
	Password       string    `json:"-"` // for disable show on return model
	Identification string    `json:"identification"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
