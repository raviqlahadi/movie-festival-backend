package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/raviqlahadi/movie-festival-backend/internal/repositories"
)

type VoteHandler struct {
	VoteRepo *repositories.VoteRepository
}

// Vote for a movie
func (h *VoteHandler) VoteMovie(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(uint)
	vars := mux.Vars(r)
	movieID, err := strconv.Atoi(vars["movie_id"])
	if err != nil {
		http.Error(w, "Invalid movie ID", http.StatusBadRequest)
		return
	}

	// Check if the user already voted
	voted, err := h.VoteRepo.HasUserVoted(userID, uint(movieID))
	if err != nil {
		http.Error(w, "Failed to check vote status", http.StatusInternalServerError)
		return
	}
	if voted {
		http.Error(w, "User has already voted for this movie", http.StatusConflict)
		return
	}

	// Add the vote
	if err := h.VoteRepo.AddVote(userID, uint(movieID)); err != nil {
		http.Error(w, "Failed to vote for movie", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Vote recorded successfully"})
}

// Unvote a movie
func (h *VoteHandler) UnvoteMovie(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(uint)
	vars := mux.Vars(r)
	movieID, err := strconv.Atoi(vars["movie_id"])
	if err != nil {
		http.Error(w, "Invalid movie ID", http.StatusBadRequest)
		return
	}

	// Remove the vote
	if err := h.VoteRepo.RemoveVote(userID, uint(movieID)); err != nil {
		http.Error(w, "Failed to remove vote", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Vote removed successfully"})
}

// List all voted movies for a user
func (h *VoteHandler) ListUserVotes(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(uint)

	movies, err := h.VoteRepo.GetUserVotes(userID)
	if err != nil {
		http.Error(w, "Failed to fetch user votes", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(movies)
}

// Get the most voted movie (admin only)
func (h *VoteHandler) MostVotedMovie(w http.ResponseWriter, r *http.Request) {
	isAdmin := r.Context().Value("is_admin").(bool)
	if !isAdmin {
		http.Error(w, "Forbidden: Admin access required", http.StatusForbidden)
		return
	}

	// Get query parameters
	query := r.URL.Query()
	page, err := strconv.Atoi(query.Get("page"))
	if err != nil || page <= 0 {
		page = 1 // Default page
	}
	limit, err := strconv.Atoi(query.Get("limit"))
	if err != nil || limit <= 0 {
		limit = 10 // Default limit
	}

	movies, err := h.VoteRepo.GetMostVotedMovies(page, limit)
	if err != nil {
		http.Error(w, "Failed to fetch the most voted movies", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"movies": movies,
		"page":   page,
		"limit":  limit,
	})
}
