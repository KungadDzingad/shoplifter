package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Mail     string `json:"mail" gorm:"primaryKey"`
	Username string `json:"username" `
	Password string `json:"password"`
}
