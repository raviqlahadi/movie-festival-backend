package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/raviqlahadi/movie-festival-backend/internal/repositories"
)

type MovieHandler struct {
	MovieRepo *repositories.MovieRepository
}

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
