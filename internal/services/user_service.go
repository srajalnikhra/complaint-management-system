package services

import (
	"github.com/srajalnikhra/complaint-management-system/internal/models"
	"github.com/srajalnikhra/complaint-management-system/internal/repositories"
	"github.com/srajalnikhra/complaint-management-system/internal/utils"
)

// UserService encapsulates business operations relating to active users.
type UserService struct {
	repo *repositories.UserRepository
}

// NewUserService creates a new instance of the UserService.
func NewUserService() *UserService {
	return &UserService{
		repo: repositories.NewUserRepository(),
	}
}

// CreateUser registers a new user in the application after securing their password.
func (s *UserService) CreateUser(user *models.User) error {

	// Securely hash the user's raw password before inserting it.
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hashedPassword

	// Save the new user record in the database.
	return s.repo.Create(user)
}
