package utils

import (
	"encoding/json"
	"net/http"
)

// APIResponse normalizes outgoing JSON data into a uniform standard.
type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// Success writes a standardized successful JSON HTTP response payload.
func Success(w http.ResponseWriter, statusCode int, message string, data interface{}) {

	// Set Response content-type metadata header to JSON configuration.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	// Encode success response values model container.
	json.NewEncoder(w).Encode(APIResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// Error writes a standardized failure JSON HTTP response payload.
func Error(w http.ResponseWriter, statusCode int, message string) {

	// Set Response content-type metadata header to JSON configuration.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	// Encode unsuccessful response values model container.
	json.NewEncoder(w).Encode(APIResponse{
		Success: false,
		Message: message,
	})
}
