package models

import (
	"gorm.io/gorm"
)

type Link struct {
	gorm.Model
	UserID    uint   `json:"user_id"`
	portfolio string `json:"portfolio"`
	twitter   string `json:"twitter"`
	github    string `json:"github"`
}
