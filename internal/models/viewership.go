package models

import "time"

type Viewership struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	MovieID   int  `json:"movie_id" gorm:"not null"`
	UserID    *int `json:"user_id"` // Nullable
	WatchTime int  `json:"watch_time" gorm:"not null"`
}

func (Viewership) TableName() string {
	return "viewership"
}
