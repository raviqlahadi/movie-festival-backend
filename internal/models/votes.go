package models

import "time"

type Vote struct {
	ID        uint      `gorm:"primaryKey"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	UserID    uint      `json:"user_id" gorm:"not null"`
	MovieID   uint      `json:"movie_id" gorm:"not null"`
}
