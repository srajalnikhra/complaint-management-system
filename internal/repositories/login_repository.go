package repositories

import (
	"context"
	"time"

	"github.com/srajalnikhra/complaint-management-system/internal/database"
	"github.com/srajalnikhra/complaint-management-system/internal/models"
)

type LoginRepository struct{}

func NewLoginRepository() *LoginRepository {
	return &LoginRepository{}
}

func (r *LoginRepository) GetByEmail(email string) (*models.User, error) {
	user := &models.User{}

	query := `
	SELECT id, name, email, password, role, is_active, created_at, updated_at
	FROM users
	WHERE email = $1;
	`

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

func (r *LoginRepository) SaveOTP(email, otpHash string, expiresAt time.Time) error {

	query := `
	UPDATE users
	SET otp_hash = $1,
	    otp_expires_at = $2,
	    updated_at = NOW()
	WHERE email = $3;
	`

	_, err := database.DB.Exec(
		context.Background(),
		query,
		otpHash,
		expiresAt,
		email,
	)

	return err
}

func (r *LoginRepository) ClearOTP(email string) error {

	query := `
	UPDATE users
	SET otp_hash = NULL,
	    otp_expires_at = NULL,
	    updated_at = NOW()
	WHERE email = $1;
	`

	_, err := database.DB.Exec(
		context.Background(),
		query,
		email,
	)

	return err
}

func (r *LoginRepository) UpdatePassword(email, hashedPassword string) error {

	query := `
	UPDATE users
	SET password = $1,
	    updated_at = NOW()
	WHERE email = $2;
	`

	_, err := database.DB.Exec(
		context.Background(),
		query,
		hashedPassword,
		email,
	)

	return err
}

