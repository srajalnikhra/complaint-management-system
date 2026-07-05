package repositories

import (
	"context"

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