package middleware

import (
	"net/http"

	"github.com/srajalnikhra/complaint-management-system/internal/models"
	"github.com/srajalnikhra/complaint-management-system/internal/utils"
)

// AdminMiddleware ensures that the request context has the admin role before allowing access.
func AdminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Retrieve the role from the request context.
		role, ok := r.Context().Value("role").(string)
		if !ok {
			utils.Error(
				w,
				http.StatusUnauthorized,
				"Unauthorized",
			)
			return
		}

		// Check if the user is an admin.
		if role != models.RoleAdmin {
			utils.Error(
				w,
				http.StatusForbidden,
				"Forbidden",
			)
			return
		}

		// Proceed to the next handler.
		next.ServeHTTP(w, r)
	})
}
