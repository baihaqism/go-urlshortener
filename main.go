package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/baihaqism/urlshortener/handlers"
	_ "github.com/lib/pq"
)

func main() {
	// Get the database connection string from environment variables
	dbConnStr := os.Getenv("DB_CONNECTION_STRING")

	// Connect to the database
	db, err := sql.Open("postgres", dbConnStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create the router and register the request handlers
	http.HandleFunc("/shorten", handlers.ShortenHandler(db))
	http.HandleFunc("/", handlers.RedirectHandler(db))

	// Start the server
	log.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
