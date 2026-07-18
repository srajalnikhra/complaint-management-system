package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/srajalnikhra/complaint-management-system/internal/repositories"
	"github.com/srajalnikhra/complaint-management-system/internal/utils"
)

// AuthMiddleware validates the JWT token in the Authorization header.
// If the token is valid and the user is active, it adds the userID and role
// to the request context.
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Retrieve the Authorization header.
		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			utils.Error(
				w,
				http.StatusUnauthorized,
				"Missing token",
			)
			return
		}

		// Extract the JWT token.
		token := strings.TrimPrefix(authHeader, "Bearer ")

		// Parse and validate the token.
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

		// Check if the user still exists in the database.
		user, err := userRepo.GetByID(claims.UserID)
		if err != nil {
			utils.Error(
				w,
				http.StatusUnauthorized,
				"Invalid user",
			)
			return
		}

		// Check if the user's account has been deactivated.
		if !user.IsActive {
			utils.Error(
				w,
				http.StatusUnauthorized,
				"Your account has been deactivated",
			)
			return
		}

		// Store the userID and role in the request context.
		ctx := context.WithValue(r.Context(), "userID", claims.UserID)

		ctx = context.WithValue(ctx, "role", claims.Role)

		// Call the next handler with the authenticated context.
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
