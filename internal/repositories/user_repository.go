package repositories

import (
	"context"
	"strings"

	"github.com/srajalnikhra/complaint-management-system/internal/database"
	"github.com/srajalnikhra/complaint-management-system/internal/models"
	"github.com/srajalnikhra/complaint-management-system/internal/utils"
)

// UserRepository handles the database operations for user profiling and creation.
type UserRepository struct{}

// NewUserRepository creates a new instance of the UserRepository.
func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

// Create inserts a new user record into the database.
func (r *UserRepository) Create(user *models.User) error {

	query := `
	INSERT INTO users (name, email, password, role)
	VALUES ($1, $2, $3, $4)
	RETURNING id, created_at, updated_at;
	`

	// Insert the user details and check for unique email violations.
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

// GetByID finds a user record by their unique user ID.
func (r *UserRepository) GetByID(id int) (*models.User, error) {

	user := &models.User{}

	query := `
	SELECT
		id,
		name,
		email,
		password,
		role,
		is_active,
		created_at,
		updated_at
	FROM users
	WHERE id = $1;
	`

	// Query SQL row for user data by ID.
	err := database.DB.QueryRow(
		context.Background(),
		query,
		id,
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
