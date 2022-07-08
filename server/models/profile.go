package models

import (
	"time"
	//
	"gorm.io/gorm"
)

type Profile struct {
	gorm.Model
	UserID     int       `json:"user_id"`
	Email      string    `json:"email"`
	Birth      time.Time `json:"birthday"`
	Gender     string    `json:"gender"`
	Address    string    `json:"address"`
	Occupation string    `json:"occupation"`
}
