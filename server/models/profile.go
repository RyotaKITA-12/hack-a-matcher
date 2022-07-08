package models

import (
	"time"
 //
	"gorm.io/gorm"
)

type Profile struct {
	gorm.Model
	UserID     int      `json:"user_id"`
	Email      string    `json:"email"`
	Birth      time.Time `json:"birthday"`
	Address    string    `json:"address"`
	Occupation string    `json:"occupation"`
}
