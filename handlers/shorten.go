package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/baihaqism/go-urlshortener/models"
	"github.com/baihaqism/go-urlshortener/utils"
)

// ShortenHandler handles requests to generate short URLs from long URLs.
func ShortenHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request models.ShortenRequest

		// Decode the request JSON body
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Generate a random short URL
		shortURL := utils.GenerateRandomString(8)

		// Insert the short URL and long URL into the database
		_, err = db.Exec("INSERT INTO urls (short_url, long_url) VALUES ($1, $2)", shortURL, request.LongURL)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Create the response object
		response := models.ShortenResponse{
			ShortURL: shortURL,
		}

		// Encode the response as JSON and send it
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
