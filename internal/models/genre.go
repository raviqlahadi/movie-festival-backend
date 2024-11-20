package models

import "time"

type Genre struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string  `json:"name" gorm:"unique;not null"`
	Movies    []Movie `json:"movies" gorm:"many2many:movie_genres"`
}
