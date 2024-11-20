package routes

import (
	"os"

	"github.com/gorilla/mux"
	"github.com/raviqlahadi/movie-festival-backend/internal/handlers"
	"github.com/raviqlahadi/movie-festival-backend/internal/middleware"
	"github.com/raviqlahadi/movie-festival-backend/internal/repositories"
	"github.com/raviqlahadi/movie-festival-backend/internal/services"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter()
	secretKey := os.Getenv("SECRET_KEY")

	// Load repositories
	genreRepo := repositories.NewGenreRepository()
	movieRepo := repositories.NewMovieRepository(genreRepo)
	userRepo := repositories.NewUserRepository()
	voteRepo := repositories.NewVoteRepository()
	authService := services.NewAuthService(secretKey)

	// Load handlers
	AdminHandler := handlers.AdminHandler{MovieRepo: movieRepo}
	MovieHandler := handlers.MovieHandler{MovieRepo: movieRepo}
	AuthHandler := handlers.AuthHandler{UserRepo: userRepo, AuthService: authService}
	VoteHandler := handlers.VoteHandler{VoteRepo: voteRepo}

	// Middleware
	authMiddleware := middleware.AuthMiddleware(secretKey)

	// Public routes
	router.HandleFunc("/auth/register", AuthHandler.Register).Methods("POST")
	router.HandleFunc("/auth/login", AuthHandler.Login).Methods("POST")
	router.HandleFunc("/movies", MovieHandler.ListMovies).Methods("GET")
	router.HandleFunc("/movies/search", MovieHandler.SearchMovies).Methods("GET")
	router.HandleFunc("/movies/{id}/view", MovieHandler.TrackViewership).Methods("POST")

	// Protected user routes
	userRouter := router.PathPrefix("/").Subrouter()
	userRouter.Use(authMiddleware) // Apply auth middleware to user routes
	userRouter.HandleFunc("/movies/{movie_id}/vote", VoteHandler.VoteMovie).Methods("POST")
	userRouter.HandleFunc("/movies/{movie_id}/unvote", VoteHandler.UnvoteMovie).Methods("DELETE")
	userRouter.HandleFunc("/user/votes", VoteHandler.ListUserVotes).Methods("GET")

	// Protected admin routes
	adminRouter := router.PathPrefix("/admin").Subrouter()
	adminRouter.Use(authMiddleware) // Apply auth middleware to admin routes
	adminRouter.HandleFunc("/movies", AdminHandler.CreateMovie).Methods("POST")
	adminRouter.HandleFunc("/movies/{id}", AdminHandler.UpdateMovie).Methods("PUT")
	adminRouter.HandleFunc("/movies/most-viewed", AdminHandler.GetMostViewedMoviesAndGenreas).Methods("GET")
	adminRouter.HandleFunc("/movies/most-voted", VoteHandler.MostVotedMovie).Methods("GET")

	return router
}
