package middleware

import (
	"net/http"

	"github.com/srajalnikhra/complaint-management-system/internal/utils"
)

// UserMiddleware ensures that the request context has the user role before allowing access.
func UserMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Get the role from the request context.
		role, ok := r.Context().Value("role").(string)

		// Check if the user is a standard user.
		if !ok || role != "user" {
			utils.Error(
				w,
				http.StatusForbidden,
				"Access denied",
			)
			return
		}

		// Proceed to the next handler.
		next.ServeHTTP(w, r)
	})
}
