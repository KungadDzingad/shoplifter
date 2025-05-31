package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Mail     string `json:"mail" gorm:"primaryKey"`
	Username string `json:"username" `
	Password string `json:"password,omitempty"`

	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
