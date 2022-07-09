package models

import (
	"gorm.io/gorm"
	"time"
)

type Profile struct {
	gorm.Model
	UserID  int       `json:"user_id"`
	Email   string    `json:"email"`
	Birth   time.Time `json:"birthday"`
	Gender  string    `json:"gender"`
	Address string    `json:"address"`
}

type Link struct {
	gorm.Model
	UserID    int    `json:"user_id"`
	Portfolio string `json:"portfolio"`
	Twitter   string `json:"twitter"`
	Github    string `json:"github"`
}

type Role struct {
	gorm.Model
	UserID uint   `json:"user_id"`
	Role   string `json:"role"`
}

type Skill struct {
	gorm.Model
	UserID     int    `json:"user_id"`
	Experience string `json:"experience"`
	Skill      string `json:"skill"`
}
