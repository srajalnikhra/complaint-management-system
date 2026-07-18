package repositories

import (
	"context"
	"time"

	"github.com/srajalnikhra/complaint-management-system/internal/database"
	"github.com/srajalnikhra/complaint-management-system/internal/models"
)

// LoginRepository performs direct database access for login checking and token issuance.
type LoginRepository struct{}

// NewLoginRepository creates a new instance of the LoginRepository.
func NewLoginRepository() *LoginRepository {
	return &LoginRepository{}
}

// GetByEmail finds a user record matching a specific email address.
func (r *LoginRepository) GetByEmail(email string) (*models.User, error) {
	user := &models.User{}

	query := `
	SELECT id, name, email, password, role, is_active, created_at, updated_at
	FROM users
	WHERE email = $1;
	`

	// Query database for a user matching the provided email.
	err := database.DB.QueryRow(
		context.Background(),
		query,
		email,
	).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.Role,
		&user.IsActive,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}

// GetOTPData gets the verification OTP hash and expiration time for a user.
func (r *LoginRepository) GetOTPData(email string) (*models.User, error) {

	user := &models.User{}

	query := `
	SELECT
		name,
		email,
		otp_hash,
		otp_expires_at
	FROM users
	WHERE email = $1;
	`

	// Query database for the OTP hash and its expiry timestamp.
	err := database.DB.QueryRow(
		context.Background(),
		query,
		email,
	).Scan(
		&user.Name,
		&user.Email,
		&user.OTPHash,
		&user.OTPExpiresAt,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}

// SaveOTP stores the generated OTP details for a user.
func (r *LoginRepository) SaveOTP(email, otpHash string, expiresAt time.Time) error {

	query := `
	UPDATE users
	SET otp_hash = $1,
	    otp_expires_at = $2,
	    updated_at = NOW()
	WHERE email = $3;
	`

	// Update sql record sets OTP hash and expires_at fields.
	_, err := database.DB.Exec(
		context.Background(),
		query,
		otpHash,
		expiresAt,
		email,
	)

	return err
}

// ClearOTP clears the OTP details from the user record.
func (r *LoginRepository) ClearOTP(email string) error {

	query := `
	UPDATE users
	SET otp_hash = NULL,
	    otp_expires_at = NULL,
	    updated_at = NOW()
	WHERE email = $1;
	`

	// Update sql record sets OTP fields to null.
	_, err := database.DB.Exec(
		context.Background(),
		query,
		email,
	)

	return err
}

// UpdatePassword updates a user's password with a new hash.
func (r *LoginRepository) UpdatePassword(email, hashedPassword string) error {

	query := `
	UPDATE users
	SET password = $1,
	    updated_at = NOW()
	WHERE email = $2;
	`

	// Save the newly hashed password in the user's database row.
	_, err := database.DB.Exec(
		context.Background(),
		query,
		hashedPassword,
		email,
	)

	return err
}
