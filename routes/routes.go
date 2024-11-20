package routes

import (
	"github.com/gorilla/mux"
	"github.com/raviqlahadi/movie-festival-backend/internal/handlers"
	"github.com/raviqlahadi/movie-festival-backend/internal/repositories"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter()

	//Inject
	movieRepo := &repositories.MovieRepository{}
	AdminHandler := handlers.AdminHandler{MovieRepo: movieRepo}

	//router path
	router.HandleFunc("/admin/movies", AdminHandler.CreateMovie).Methods("POST")
	router.HandleFunc("/admin/movies/{id}", AdminHandler.UpdateMovie).Methods("PUT")

	return router
}
