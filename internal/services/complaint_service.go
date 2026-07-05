package services

import (
	"github.com/srajalnikhra/complaint-management-system/internal/models"
	"github.com/srajalnikhra/complaint-management-system/internal/repositories"
)

type ComplaintService struct {
	repo *repositories.ComplaintRepository
}

func NewComplaintService() *ComplaintService {
	return &ComplaintService{
		repo: repositories.NewComplaintRepository(),
	}
}

func (s *ComplaintService) Create(c *models.Complaint) error {
	return s.repo.Create(c)
}

func (s *ComplaintService) GetByUserID(userID int) ([]models.Complaint, error) {
	return s.repo.GetByUserID(userID)
}

func (s *ComplaintService) GetByID(id int) (*models.Complaint, error) {
	return s.repo.GetByID(id)
}

func (s *ComplaintService) Update(c *models.Complaint) error {
	return s.repo.Update(c)
}

func (s *ComplaintService) Delete(id int) error {
	return s.repo.Delete(id)
}