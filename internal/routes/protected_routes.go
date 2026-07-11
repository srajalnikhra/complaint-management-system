package routes

import (
	"net/http"

	"github.com/srajalnikhra/complaint-management-system/internal/middleware"
)

func RegisterProtectedRoutes() {

	http.Handle(
		"/profile",
		middleware.AuthMiddleware(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("Protected Route"))
			}),
		),
	)
}
