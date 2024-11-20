package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/raviqlahadi/movie-festival-backend/internal/models"
	"github.com/raviqlahadi/movie-festival-backend/internal/repositories"
)

type MovieHandler struct {
	MovieRepo *repositories.MovieRepository
}

// Get movies with pagination
func (h *MovieHandler) ListMovies(w http.ResponseWriter, r *http.Request) {
	//Get query parameter for paginations
	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")

	//Set default
	page := 1
	limit := 10

	//parse page and limit
	if pageStr != "" {
		if parsedPage, err := strconv.Atoi(pageStr); err == nil && parsedPage > 0 {
			page = parsedPage
		}
	}

	if limitStr != "" {
		if parseLimit, err := strconv.Atoi(limitStr); err == nil && parseLimit > 0 {
			limit = parseLimit
		}
	}

	offset := (page - 1) * limit
	movies, err := h.MovieRepo.GetMoviesWithPagination(offset, limit)
	if err != nil {
		http.Error(w, "Failed to fetch movies", http.StatusInternalServerError)
		return
	}

	// Respond with paginated results
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"page":   page,
		"limit":  limit,
		"movies": movies,
	})

}

// Search movies by Keyword or Genre
func (h *MovieHandler) SearchMovies(w http.ResponseWriter, r *http.Request) {
	// Get query parameters
	query := r.URL.Query().Get("query")
	genre := r.URL.Query().Get("genre")
	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")

	// Validate inputs
	if query == "" && genre == "" {
		http.Error(w, "Query or genre parameter is required", http.StatusBadRequest)
		return
	}

	page := 1
	limit := 10

	if pageStr != "" {
		if parsedPage, err := strconv.Atoi(pageStr); err == nil && parsedPage > 0 {
			page = parsedPage
		}
	}

	if limitStr != "" {
		if parsedLimit, err := strconv.Atoi(limitStr); err == nil && parsedLimit > 0 && parsedLimit <= 100 {
			limit = parsedLimit
		}
	}

	offset := (page - 1) * limit

	var movies []models.Movie
	var err error

	// Perform search based on query or genre
	if genre != "" {
		movies, err = h.MovieRepo.SearchMoviesByGenre(genre, offset, limit)
	} else {
		movies, err = h.MovieRepo.SearchMovies(query, offset, limit)
	}

	if err != nil {
		http.Error(w, "Failed to search movies", http.StatusInternalServerError)
		return
	}

	// Respond with the search results
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"page":   page,
		"limit":  limit,
		"movies": movies,
	})
}
