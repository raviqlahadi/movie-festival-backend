package main

import (
	"log"
	"net/http"

	"github.com/raviqlahadi/movie-festival-backend/internal/db"
)

func main() {
	//Init the database connection
	db.ConnectDB()

	//Set up routes

	//Start the server
	log.Println("Server is running on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}

}
