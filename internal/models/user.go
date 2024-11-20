package models

import "time"

type User struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Username  string `json:"username" gorm:"unique;not null"`
	Email     string `json:"email" gorm:"unique;not null"`
	Password  string `json:"-" gorm:"not null"` // Excluded from JSON responses
	IsAdmin   bool   `json:"is_admin" gorm:"default:false"`
}
