package services

import (
	"github.com/srajalnikhra/complaint-management-system/internal/models"
	"github.com/srajalnikhra/complaint-management-system/internal/repositories"
)

type AdminService struct {
	repo *repositories.AdminRepository
}

func NewAdminService() *AdminService {
	return &AdminService{
		repo: repositories.NewAdminRepository(),
	}
}

func (s *AdminService) GetAllComplaints(page, limit int, search string) ([]models.Complaint, error) {
	return s.repo.GetAllComplaints(page, limit, search)
}

func (s *AdminService) UpdateComplaintStatus(id int, status string) error {
	return s.repo.UpdateComplaintStatus(id, status)
}
