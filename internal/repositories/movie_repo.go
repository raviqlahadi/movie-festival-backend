package repositories

import (
	"github.com/raviqlahadi/movie-festival-backend/internal/db"
	"github.com/raviqlahadi/movie-festival-backend/internal/models"
)

type MovieRepository struct {
	GenreRepo *GenreRepository
}

func NewMovieRepository(genreRepo *GenreRepository) *MovieRepository {
	return &MovieRepository{GenreRepo: genreRepo}
}

// Adding new movie record
func (r *MovieRepository) CreateMovieWithGenres(movie models.Movie, genres []string) error {
	tx := db.DB.Begin()

	//Crete the movie record
	if err := db.DB.Create(&movie).Error; err != nil {
		tx.Rollback()
		return err
	}

	//Handle Genres
	for _, genreName := range genres {
		genreId, err := r.GenreRepo.GetOrCreateGenreID(genreName)
		if err != nil {
			tx.Rollback()
			return err
		}

		//Insert into movie_genres table
		if err := tx.Exec(`
			INSERT INTO 
				movie_genres (movie_id, genre_id)
			VALUES (?, ?) ON CONFLICT DO NOTHING 
		`, movie.ID, genreId).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}

// Update movie by id
func (r *MovieRepository) UpdateMovieWithGenres(id int, updatedMovie models.Movie, genres []string) error {
	// Start a transaction
	tx := db.DB.Begin()

	// Update the movie
	if err := tx.Model(&models.Movie{}).Where("id = ?", id).Updates(updatedMovie).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Clear existing genres for the movie
	if err := tx.Exec(`
		DELETE FROM movie_genres WHERE movie_id = ?
	`, id).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Handle genres
	for _, genreName := range genres {
		genreID, err := r.GenreRepo.GetOrCreateGenreID(genreName)
		if err != nil {
			tx.Rollback()
			return err
		}

		// Insert into movie_genres table
		if err := tx.Exec(`
			INSERT INTO movie_genres (movie_id, genre_id)
			VALUES (?, ?) ON CONFLICT DO NOTHING
		`, id, genreID).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	// Commit the transaction
	return tx.Commit().Error
}

// Get most viewed movie
func (r *MovieRepository) GetMostViewedMovies(limit int) ([]models.Movie, error) {
	if limit <= 0 {
		limit = 5
	}
	if limit > 100 {
		limit = 100
	}
	var movie []models.Movie
	err := db.DB.Order("view_count desc").Limit(limit).Find(&movie).Error
	return movie, err
}

// Get most viewed genre
type GenreViewCount struct {
	Name      string
	ViewCount int
}

func (r *MovieRepository) GetMostViewedGenre(limit int) ([]GenreViewCount, error) {
	if limit <= 0 {
		limit = 5
	}
	if limit > 100 {
		limit = 100
	}
	var genres []GenreViewCount

	err := db.DB.Raw(`
		SELECT 
			g.name,
			SUM(m.view_count) as view_count 
		FROM 
			genres g
		JOIN 
			movie_genres mg ON g.id = mg.genre_id 
		JOIN 
			movies m ON mg.movie_id = m.id 
		GROUP BY
			 g.name
		ORDER BY 
			SUM(m.view_count) DESC
		LIMIT ?
	`, limit).Scan(&genres).Error
	return genres, err
}

// Get movies with pagination
func (r *MovieRepository) GetMoviesWithPagination(offset, limit int) ([]models.Movie, error) {
	var movies []models.Movie
	err := db.DB.Preload("Genres").Offset(offset).Limit(limit).Find(&movies).Error
	return movies, err
}
