package models

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	UserID     int       `json:"user_id"`
	Title      string    `json:"title"`
	RecruitNum int       `json:"recruit_num"`
	Content    string    `json:"content"`
	ViewPeriod time.Time `json:"view_period"`
}

type PostRole struct {
	gorm.Model
	PostID int    `json:"post_id"`
	Role   string `json:"role"`
}
