package routes

import (
	"github.com/gorilla/mux"
	"github.com/raviqlahadi/movie-festival-backend/internal/handlers"
	"github.com/raviqlahadi/movie-festival-backend/internal/repositories"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter()

	//Inject
	genreRepo := repositories.NewGenreRepository()
	movieRepo := repositories.NewMovieRepository(genreRepo)

	AdminHandler := handlers.AdminHandler{MovieRepo: movieRepo}

	//router path
	router.HandleFunc("/admin/movies", AdminHandler.CreateMovie).Methods("POST")
	router.HandleFunc("/admin/movies/{id}", AdminHandler.UpdateMovie).Methods("PUT")
	router.HandleFunc("/admin/movies/most-viewed", AdminHandler.GetMostViewedMoviesAndGenreas).Methods("GET")

	return router
}
