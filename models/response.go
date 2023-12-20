package utils

import (
	"encoding/json"
	"net/http"
)

// ErrorResponse represents an error response.
type ErrorResponse struct {
	Message string `json:"message"`
}

// SendErrorResponse sends an error response with the specified message.
func SendErrorResponse(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := ErrorResponse{
		Message: message,
	}

	json.NewEncoder(w).Encode(response)
}
