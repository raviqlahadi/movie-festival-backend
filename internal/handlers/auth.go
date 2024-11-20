package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/raviqlahadi/movie-festival-backend/internal/models"
	"github.com/raviqlahadi/movie-festival-backend/internal/repositories"
	"github.com/raviqlahadi/movie-festival-backend/internal/services"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	UserRepo    *repositories.UserRepository
	AuthService *services.AuthService
}

func NewAuthHandler(userRepo *repositories.UserRepository, authService *services.AuthService) *AuthHandler {
	return &AuthHandler{
		UserRepo:    userRepo,
		AuthService: authService,
	}
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"passowrd"`
		IsAdmin  bool   `json:"is_admin"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	hashedPasswor, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Password hash failed", http.StatusInternalServerError)
		return
	}

	user := models.User{
		Username: request.Username,
		Email:    request.Email,
		Password: string(hashedPasswor),
		IsAdmin:  request.IsAdmin,
	}
	if err := h.UserRepo.CreateUser(user); err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	user, err := h.UserRepo.GetUserByEmail(request.Email)
	if err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	token, err := h.AuthService.GenerateToken(user.ID, user.IsAdmin)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
