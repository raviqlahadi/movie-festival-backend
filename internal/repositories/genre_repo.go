package repositories

import (
	"github.com/raviqlahadi/movie-festival-backend/internal/db"
	"github.com/raviqlahadi/movie-festival-backend/internal/models"
)

type GenreRepository struct{}

func NewGenreRepository() *GenreRepository {
	return &GenreRepository{}
}

// Get or Created genre by name if not exist
func (r *GenreRepository) GetOrCreateGenreID(genreName string) (int, error) {
	var genre models.Genre

	// Check if the genre exists
	err := db.DB.Where("name = ?", genreName).First(&genre).Error
	if err == nil {
		return int(genre.ID), nil
	}

	// If genre not found, create it
	genre = models.Genre{Name: genreName}
	if err := db.DB.Create(&genre).Error; err != nil {
		return 0, err
	}
	return int(genre.ID), nil
}
