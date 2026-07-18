package routes

import (
	"net/http"

	"github.com/srajalnikhra/complaint-management-system/internal/middleware"
)

// RegisterProtectedRoutes defines general authenticated endpoints.
func RegisterProtectedRoutes() {

	// Map profile access wrapped in the AuthMiddleware.
	http.Handle(
		"/profile",
		middleware.AuthMiddleware(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("Protected Route"))
			}),
		),
	)
}
