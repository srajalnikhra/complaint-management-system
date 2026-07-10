package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/srajalnikhra/complaint-management-system/internal/repositories"
	"github.com/srajalnikhra/complaint-management-system/internal/utils"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			utils.Error(
				w,
				http.StatusUnauthorized,
				"Missing token",
			)
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")

		claims, err := utils.ValidateJWT(token)
		if err != nil {
			utils.Error(
				w,
				http.StatusUnauthorized,
				"Invalid token",
			)
			return
		}

		userRepo := repositories.NewUserRepository()

		user, err := userRepo.GetByID(claims.UserID)
		if err != nil {
			utils.Error(
				w,
				http.StatusUnauthorized,
				"Invalid user",
			)
			return
		}

		if !user.IsActive {
			utils.Error(
				w,
				http.StatusUnauthorized,
				"Your account has been deactivated",
			)
			return
		}

		ctx := context.WithValue(r.Context(), "userID", claims.UserID)

		ctx = context.WithValue(ctx, "role", claims.Role)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
