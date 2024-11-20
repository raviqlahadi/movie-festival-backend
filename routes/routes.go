package routes

import (
	"os"

	"github.com/gorilla/mux"
	"github.com/raviqlahadi/movie-festival-backend/internal/handlers"
	"github.com/raviqlahadi/movie-festival-backend/internal/repositories"
	"github.com/raviqlahadi/movie-festival-backend/internal/services"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter()
	secret_key := os.Getenv("SECRET_KEY")

	//Load repository
	genreRepo := repositories.NewGenreRepository()
	movieRepo := repositories.NewMovieRepository(genreRepo)
	userRepo := repositories.NewUserRepository()
	voteRepo := repositories.NewVoteRepository()
	authService := services.NewAuthService(secret_key)

	AdminHandler := handlers.AdminHandler{MovieRepo: movieRepo}
	MovieHandler := handlers.MovieHandler{MovieRepo: movieRepo}
	AuthHandler := handlers.AuthHandler{UserRepo: userRepo, AuthService: authService}
	VoteHandler := handlers.VoteHandler{VoteRepo: voteRepo}

	//router path
	router.HandleFunc("/admin/movies", AdminHandler.CreateMovie).Methods("POST")
	router.HandleFunc("/admin/movies/{id}", AdminHandler.UpdateMovie).Methods("PUT")
	router.HandleFunc("/admin/movies/most-viewed", AdminHandler.GetMostViewedMoviesAndGenreas).Methods("GET")
	router.HandleFunc("/admin/movies/most-voted", VoteHandler.MostVotedMovie).Methods("GET")
	router.HandleFunc("/auth/register", AuthHandler.Register).Methods("POST")
	router.HandleFunc("/auth/login", AuthHandler.Login).Methods("POST")
	router.HandleFunc("/movies", MovieHandler.ListMovies).Methods("GET")
	router.HandleFunc("/movies/search", MovieHandler.SearchMovies).Methods("GET")
	router.HandleFunc("/movies/{id}/view", MovieHandler.TrackViewership).Methods("POST")
	router.HandleFunc("/movies/{movie_id}/vote", VoteHandler.VoteMovie).Methods("POST")
	router.HandleFunc("/movies/{movie_id}/unvote", VoteHandler.UnvoteMovie).Methods("DELETE")
	router.HandleFunc("/user/votes", VoteHandler.ListUserVotes).Methods("GET")

	return router
}
