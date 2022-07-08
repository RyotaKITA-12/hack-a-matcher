package models

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	UserID uint   `json:"user_id"`
	Role   string `json:"role"`
}
