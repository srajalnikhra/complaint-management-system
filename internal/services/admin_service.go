package services

import (
	"github.com/srajalnikhra/complaint-management-system/internal/dto"
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

func (s *AdminService) GetAllComplaints(page, limit int, search, status, sort, order string) ([]models.Complaint, error) {
	return s.repo.GetAllComplaints(page, limit, search, status, sort, order)
}

func (s *AdminService) UpdateComplaintStatus(id int, status string) error {

	err := s.repo.UpdateComplaintStatus(id, status)
	if err != nil {
		return err
	}

	data, err := s.repo.GetComplaintEmailData(id)
	if err != nil {
		return nil
	}

	err = SendComplaintStatusEmail(
		data.UserEmail,
		data.UserName,
		data.ComplaintTitle,
		data.Status,
	)

	if err != nil {
		return nil
	}

	return nil
}

func (s *AdminService) GetAllUsers() ([]dto.AdminUserResponse, error) {
	return s.repo.GetAllUsers()
}

func (s *AdminService) UpdateUserRole(id int, role string) error {
	return s.repo.UpdateUserRole(id, role)
}

func (s *AdminService) UpdateUserStatus(id int, isActive bool) error {
	return s.repo.UpdateUserStatus(id, isActive)
}

func (s *AdminService) DeleteUser(id int) error {
	return s.repo.DeleteUser(id)
}
