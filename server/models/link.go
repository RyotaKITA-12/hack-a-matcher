package models

import (
	"gorm.io/gorm"
)

type Link struct {
	gorm.Model
	UserID    int    `json:"user_id"`
	Portfolio string `json:"portfolio"`
	Twitter   string `json:"twitter"`
	Github    string `json:"github"`
}
