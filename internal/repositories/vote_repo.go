package repositories

import (
	"github.com/raviqlahadi/movie-festival-backend/internal/db"
	"github.com/raviqlahadi/movie-festival-backend/internal/models"
)

type VoteRepository struct{}

func NewVoteRepository() *VoteRepository {
	return &VoteRepository{}
}

// Check if a user has already voted for a movie
func (r *VoteRepository) HasUserVoted(userID, movieID uint) (bool, error) {
	var count int64
	err := db.DB.Model(&models.Vote{}).
		Where("user_id = ? AND movie_id = ?", userID, movieID).
		Count(&count).Error
	return count > 0, err
}

// Add a vote
func (r *VoteRepository) AddVote(userID, movieID uint) error {
	vote := models.Vote{
		UserID:  userID,
		MovieID: movieID,
	}
	return db.DB.Create(&vote).Error
}

// Remove a vote
func (r *VoteRepository) RemoveVote(userID, movieID uint) error {
	return db.DB.Where("user_id = ? AND movie_id = ?", userID, movieID).Delete(&models.Vote{}).Error
}

// Get all voted movies by a user
func (r *VoteRepository) GetUserVotes(userID uint) ([]models.Movie, error) {
	var movies []models.Movie
	err := db.DB.Model(&models.Movie{}).
		Select("movies.*").
		Joins("JOIN votes ON votes.movie_id = movies.id").
		Joins("LEFT JOIN movie_genres ON movie_genres.movie_id = movies.id").
		Joins("LEFT JOIN genres ON genres.id = movie_genres.genre_id").
		Where("votes.user_id = ?", userID).
		Preload("Genres").
		Find(&movies).Error

	return movies, err
}

// Get the most voted movie
func (r *VoteRepository) GetMostVotedMovies(page, limit int) ([]map[string]interface{}, error) {
	var results []map[string]interface{}

	// Calculate offset
	offset := (page - 1) * limit

	err := db.DB.Model(&models.Movie{}).
		Select("movies.*, COUNT(votes.id) AS vote_count").
		Joins("JOIN votes ON votes.movie_id = movies.id").
		Joins("LEFT JOIN movie_genres ON movie_genres.movie_id = movies.id").
		Joins("LEFT JOIN genres ON genres.id = movie_genres.genre_id").
		Group("movies.id").
		Order("vote_count DESC").
		Offset(offset).
		Limit(limit).
		Preload("Genres"). // Preload genres for each movie
		Find(&results).Error

	return results, err
}
