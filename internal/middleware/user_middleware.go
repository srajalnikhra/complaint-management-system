package middleware

import (
	"net/http"

	"github.com/srajalnikhra/complaint-management-system/internal/utils"
)

func UserMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		role, ok := r.Context().Value("role").(string)
		if !ok || role != "user" {
			utils.Error(
				w,
				http.StatusForbidden,
				"Access denied",
			)
			return
		}

		next.ServeHTTP(w, r)
	})
}
