package handlers

import (
	"database/sql"
	"net/http"
)

// RedirectHandler handles requests to redirect short URLs to long URLs.
func RedirectHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		shortURL := r.URL.Path[1:] // Remove the leading slash

		// Query the database to get the long URL for the given short URL
		var longURL string
		err := db.QueryRow("SELECT long_url FROM urls WHERE short_url = $1", shortURL).Scan(&longURL)
		if err != nil {
			http.NotFound(w, r)
			return
		}

		// Redirect to the long URL
		http.Redirect(w, r, longURL, http.StatusFound)
	}
}
