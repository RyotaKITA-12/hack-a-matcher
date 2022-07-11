package models

import "gorm.io/gorm"

type User struct {
    gorm.Model
    Name     string `json:"user_name" gorm:"unique"`
    Email    string `json:"email" gorm:"unique"`
    Password []byte `json:"-"`
}
