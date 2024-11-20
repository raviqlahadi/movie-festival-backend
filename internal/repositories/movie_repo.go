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
