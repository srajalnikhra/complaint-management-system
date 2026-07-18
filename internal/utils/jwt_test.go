package utils

import (
	"testing"

	"github.com/srajalnikhra/complaint-management-system/internal/models"
)

// TestGenerateAndValidateJWT verifies that a generated JWT token
// can be successfully validated while preserving the original claims.
func TestGenerateAndValidateJWT(t *testing.T) {
	user := &models.User{
		ID:    1,
		Email: "test@example.com",
		Role:  models.RoleUser,
	}

	token, err := GenerateJWT(user)
	if err != nil {
		t.Fatalf("failed to generate jwt: %v", err)
	}

	claims, err := ValidateJWT(token)
	if err != nil {
		t.Fatalf("failed to validate jwt: %v", err)
	}

	if claims.UserID != user.ID {
		t.Errorf("expected user id %d, got %d", user.ID, claims.UserID)
	}

	if claims.Email != user.Email {
		t.Errorf("expected email %s, got %s", user.Email, claims.Email)
	}

	if claims.Role != user.Role {
		t.Errorf("expected role %s, got %s", user.Role, claims.Role)
	}
}

// TestValidateJWTInvalidToken ensures that malformed or invalid JWT tokens
// are rejected during validation.
func TestValidateJWTInvalidToken(t *testing.T) {
	_, err := ValidateJWT("invalid.token.here")

	if err == nil {
		t.Error("expected error for invalid token")
	}
}