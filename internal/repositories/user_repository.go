package repositories

import (
	"context"
	"strings"

	"github.com/srajalnikhra/complaint-management-system/internal/database"
	"github.com/srajalnikhra/complaint-management-system/internal/models"
	"github.com/srajalnikhra/complaint-management-system/internal/utils"
)

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) Create(user *models.User) error {

	query := `
	INSERT INTO users (name, email, password, role)
	VALUES ($1, $2, $3, $4)
	RETURNING id, created_at, updated_at;
	`

	err := database.DB.QueryRow(
		context.Background(),
		query,
		user.Name,
		user.Email,
		user.Password,
		user.Role,
	).Scan(
		&user.ID,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return utils.ErrUserAlreadyExists
		}
		return err
	}

	return nil
}