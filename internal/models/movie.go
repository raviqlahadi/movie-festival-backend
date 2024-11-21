package models

import "time"

type Movie struct {
	ID          uint `gorm:"primaryKey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Title       string  `json:"title" gorm:"not null"`
	Description string  `json:"description" gorm:"not null"`
	Duration    int     `json:"duration" gorm:"not null"`
	Artists     string  `json:"artists" gorm:"not null"`
	Genres      []Genre `json:"genres" gorm:"many2many:movie_genres"`
	WatchURL    string  `json:"watch_url" gorm:"not null"`
	ViewCount   int     `json:"view_count" gorm:"default:0"`
	Year        int     `json:"year" gorm:"not null"`
}
