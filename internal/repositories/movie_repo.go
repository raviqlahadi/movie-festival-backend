package repositories

import (
	"github.com/raviqlahadi/movie-festival-backend/internal/db"
	"github.com/raviqlahadi/movie-festival-backend/internal/models"
)

type MovieRepository struct{}

// Adding new movie record
func (r *MovieRepository) CreateMovie(movie models.Movie) error {
	return db.DB.Create(&movie).Error
}

// Update movie by id
func (r *MovieRepository) UpdateMovie(id int, updatedMovie models.Movie) error {
	return db.DB.Model(&models.Movie{}).Where("id = ?", id).Updates(updatedMovie).Error
}

// Get most viewed movie
func (r *MovieRepository) GetMostViewed() (models.Movie, error) {
	var movie models.Movie
	err := db.DB.Order("view_count desc").First(&movie).Error
	return movie, err
}

// Get movies with pagination
func (r *MovieRepository) GetMoviesWithPagination(offset, limit int) ([]models.Movie, error) {
	var movies []models.Movie
	err := db.DB.Offset(offset).Limit(limit).Find(&movies).Error
	return movies, err
}
