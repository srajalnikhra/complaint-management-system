package middleware

import (
	"net/http"

	"github.com/srajalnikhra/complaint-management-system/internal/utils"
)

func AdminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		role, ok := r.Context().Value("role").(string)
		if !ok {
			utils.Error(
				w,
				http.StatusUnauthorized,
				"Unauthorized",
			)
			return
		}

		if role != "ADMIN" {
			utils.Error(
				w,
				http.StatusForbidden,
				"Forbidden",
			)
			return
		}

		next.ServeHTTP(w, r)
	})
}
