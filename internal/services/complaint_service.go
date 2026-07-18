package services

import (
	"github.com/srajalnikhra/complaint-management-system/internal/models"
	"github.com/srajalnikhra/complaint-management-system/internal/repositories"
)

// ComplaintService encapsulates the logic for user-side complaint handling.
type ComplaintService struct {
	repo *repositories.ComplaintRepository
}

// NewComplaintService creates a new instance of the ComplaintService.
func NewComplaintService() *ComplaintService {
	return &ComplaintService{
		repo: repositories.NewComplaintRepository(),
	}
}

// Create saves a new user complaint in the database.
func (s *ComplaintService) Create(c *models.Complaint) error {
	return s.repo.Create(c)
}

// GetByUserID finds all complaints submitted by a given user.
func (s *ComplaintService) GetByUserID(userID int) ([]models.Complaint, error) {
	return s.repo.GetByUserID(userID)
}

// GetByID finds a single complaint record by its ID.
func (s *ComplaintService) GetByID(id int) (*models.Complaint, error) {
	return s.repo.GetByID(id)
}

// Update edits the content details of an existing complaint.
func (s *ComplaintService) Update(c *models.Complaint) error {
	return s.repo.Update(c)
}

// Delete permanently erases a complaint from the database.
func (s *ComplaintService) Delete(id int) error {
	return s.repo.Delete(id)
}
