package models

import "gorm.io/gorm"

type User struct {
    gorm.Model
    Name     string `json:"user_name"`
    Email    string `json:"email" gorm:"unique"`
    Password []byte `json:"-"`
}
