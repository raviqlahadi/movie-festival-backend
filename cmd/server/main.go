package main

import (
	"log"
	"net/http"

	"github.com/raviqlahadi/movie-festival-backend/internal/db"
	"github.com/raviqlahadi/movie-festival-backend/routes"
)

func main() {
	//Init the database connection
	db.ConnectDB()

	//Set up routes
	router := routes.InitRoutes()
	//Start the server
	log.Println("Server is running on port 8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}

}
