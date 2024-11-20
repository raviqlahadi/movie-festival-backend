package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/raviqlahadi/movie-festival-backend/internal/models"
	"github.com/raviqlahadi/movie-festival-backend/internal/repositories"
)

type AdminHandler struct {
	MovieRepo *repositories.MovieRepository
}

func (h *AdminHandler) CreateMovie(w http.ResponseWriter, r *http.Request) {
	// Check if the user is an admin
	isAdmin := r.Context().Value("is_admin").(bool)
	if !isAdmin {
		http.Error(w, "Forbidden: Admin access required", http.StatusForbidden)
		return
	}

	var request struct {
		Movie  models.Movie `json:"movie"`
		Genres []string     `json:"genres"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := h.MovieRepo.CreateMovieWithGenres(request.Movie, request.Genres); err != nil {
		http.Error(w, "Failed to create movie with genres", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Movie created successfully"})
}

func (h *AdminHandler) UpdateMovie(w http.ResponseWriter, r *http.Request) {
	// Check if the user is an admin
	isAdmin := r.Context().Value("is_admin").(bool)
	if !isAdmin {
		http.Error(w, "Forbidden: Admin access required", http.StatusForbidden)
		return
	}

	// Extract movie ID
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid movie ID", http.StatusBadRequest)
		return
	}

	var request struct {
		Movie  models.Movie `json:"movie"`
		Genres []string     `json:"genres"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := h.MovieRepo.UpdateMovieWithGenres(id, request.Movie, request.Genres); err != nil {
		http.Error(w, "Failed to update movie with genres", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Movie updated successfully"})
}

func (h *AdminHandler) GetMostViewedMoviesAndGenreas(w http.ResponseWriter, r *http.Request) {
	// Check if the user is an admin
	isAdmin := r.Context().Value("is_admin").(bool)
	if !isAdmin {
		http.Error(w, "Forbidden: Admin access required", http.StatusForbidden)
		return
	}

	limitStr := r.URL.Query().Get("limit")
	limit := 5
	if limitStr != "" {
		parsedLimit, err := strconv.Atoi(limitStr)
		if err != nil {
			limit = parsedLimit
		}
	}
	// Get most viewed movies
	movies, err := h.MovieRepo.GetMostViewedMovies(limit)
	if err != nil {
		http.Error(w, "Failed to fetch most viewed movies", http.StatusInternalServerError)
		return
	}

	// Get most viewed genres
	genres, err := h.MovieRepo.GetMostViewedGenre(limit)
	if err != nil {
		http.Error(w, "Failed to fetch most viewed genre", http.StatusInternalServerError)
		return
	}

	//Endpoint response
	response := map[string]interface{}{
		"most_viewed_movies": movies,
		"most_viewed_genres": genres,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
