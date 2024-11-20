package models

import "time"

type Genre struct {
	ID        uint      `gorm:"primaryKey"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	Name      string    `json:"name" gorm:"unique;not null"`
	Movies    []Movie   `json:"-" gorm:"many2many:movie_genres"`
}
