package middleware

import (
	"log"
	"net/http"
	"time"
)

// LoggingMiddleware logs the HTTP method, request path, and duration for each request.
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Record the start time of the request.
		start := time.Now()

		// Run the next handler.
		next.ServeHTTP(w, r)

		// Log the request method, path, and latency.
		log.Printf(
			"%s %s %v",
			r.Method,
			r.URL.Path,
			time.Since(start),
		)
	})
}
