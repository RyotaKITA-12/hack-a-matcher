package models

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	UserID     int       `json:"user_id"`
	UserName   string    `json:"user_name"`
	Post_type  string    `json:"post_type"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	ViewPeriod time.Time `json:"view_period"`
}

type PostSkill struct {
	gorm.Model
	PostID int    `json:"post_id"`
	Lang   string `json:"lang"`
}
