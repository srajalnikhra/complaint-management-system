package services

import (
	"github.com/srajalnikhra/complaint-management-system/internal/dto"
	"github.com/srajalnikhra/complaint-management-system/internal/models"
	"github.com/srajalnikhra/complaint-management-system/internal/repositories"
)

// AdminService dictates the core logic and workflow of administrative features.
type AdminService struct {
	repo *repositories.AdminRepository
}

// NewAdminService creates a new instance of the AdminService.
func NewAdminService() *AdminService {
	return &AdminService{
		repo: repositories.NewAdminRepository(),
	}
}

// GetAllComplaints gets a list of filtered and sorted complaints.
func (s *AdminService) GetAllComplaints(page, limit int, search, status, sort, order string) ([]models.Complaint, error) {
	return s.repo.GetAllComplaints(page, limit, search, status, sort, order)
}

// UpdateComplaintStatus updates complaint status and triggers notification emails.
func (s *AdminService) UpdateComplaintStatus(id int, status string) error {

	// Update status in the database first.
	err := s.repo.UpdateComplaintStatus(id, status)
	if err != nil {
		return err
	}

	// Get complaint and user details for the notification email.
	data, err := s.repo.GetComplaintEmailData(id)
	if err != nil {
		return nil
	}

	// Send notification email to the user status changed.
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

// GetAllUsers gets all accounts registered in the application.
func (s *AdminService) GetAllUsers() ([]dto.AdminUserResponse, error) {
	return s.repo.GetAllUsers()
}

// UpdateUserRole updates the account permission role of a user.
func (s *AdminService) UpdateUserRole(id int, role string) error {
	return s.repo.UpdateUserRole(id, role)
}

// UpdateUserStatus updates the active flag status of a user's account.
func (s *AdminService) UpdateUserStatus(id int, isActive bool) error {
	return s.repo.UpdateUserStatus(id, isActive)
}

// DeleteUser deletes a specific user account entirely.
func (s *AdminService) DeleteUser(id int) error {
	return s.repo.DeleteUser(id)
}
