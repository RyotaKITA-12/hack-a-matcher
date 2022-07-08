package models

import (
	"gorm.io/gorm"
)

type Skill struct {
	gorm.Model
	UserID     int    `json:"user_id"`
	Experience string `json:"experience"`
	Skill      string `json:"skill"`
}
