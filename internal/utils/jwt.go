package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/srajalnikhra/complaint-management-system/internal/config"
	"github.com/srajalnikhra/complaint-management-system/internal/models"
)

// Claims defines the payload fields embedded within the application's JWTs.
type Claims struct {
	UserID int    `json:"user_id"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

// GenerateJWT creates a signed JSON Web Token (JWT) encapsulating user credentials.
func GenerateJWT(user *models.User) (string, error) {
	cfg := config.LoadJWTConfig()

	// Create customized claims container mapping authorization information.
	claims := Claims{
		UserID: user.ID,
		Email:  user.Email,
		Role:   user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	// Sign JWT content using standard HMAC-SHA256 protocol.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(cfg.Secret))
}

// ValidateJWT decrypts and verifies the signature of a client JWT token string.
func ValidateJWT(tokenString string) (*Claims, error) {

	// Parse the token string and reconstruct authorization claims information.
	token, err := jwt.ParseWithClaims(
		tokenString,
		&Claims{},
		func(token *jwt.Token) (interface{}, error) {
			cfg := config.LoadJWTConfig()
			return []byte(cfg.Secret), nil
		},
	)

	if err != nil {
		return nil, err
	}

	// Validate signature validity and ensure expirations parameters have not lapsed.
	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
