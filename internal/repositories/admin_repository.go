package repositories

import (
	"context"
	"fmt"

	"github.com/srajalnikhra/complaint-management-system/internal/database"
	"github.com/srajalnikhra/complaint-management-system/internal/dto"
	"github.com/srajalnikhra/complaint-management-system/internal/models"
	"github.com/srajalnikhra/complaint-management-system/internal/utils"
)

// AdminRepository provides database actions for administrative endpoints.
type AdminRepository struct{}

// NewAdminRepository creates a new instance of AdminRepository.
func NewAdminRepository() *AdminRepository {
	return &AdminRepository{}
}

// GetAllComplaints retrieves complaints list from the database based on filters and sorting options.
func (r *AdminRepository) GetAllComplaints(page, limit int, search, status, sort, order string) ([]models.Complaint, error) {

	// Calculate offset for pagination.
	offset := (page - 1) * limit

	// Allowed fields for sorting.
	allowedSorts := map[string]bool{
		"id":         true,
		"created_at": true,
		"updated_at": true,
	}

	// Validate the sort field.
	if !allowedSorts[sort] {
		sort = "id"
	}

	// Validate the sort order direction.
	if order != "ASC" && order != "DESC" {
		order = "DESC"
	}

	query := fmt.Sprintf(`
SELECT
	id,
	user_id,
	title,
	description,
	status,
	created_at,
	updated_at
FROM complaints
WHERE
(
	title ILIKE '%%' || $1 || '%%'
	OR description ILIKE '%%' || $1 || '%%'
)
AND
(
	$2 = ''
	OR status = $2
)
ORDER BY %s %s
LIMIT $3 OFFSET $4;
`, sort, order)

	// Query the database.
	rows, err := database.DB.Query(
		context.Background(),
		query,
		search,
		status,
		limit,
		offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var complaints []models.Complaint

	// Iterate and scan results into complaints.
	for rows.Next() {

		var complaint models.Complaint

		// Scan row columns into complaint structure.
		err := rows.Scan(
			&complaint.ID,
			&complaint.UserID,
			&complaint.Title,
			&complaint.Description,
			&complaint.Status,
			&complaint.CreatedAt,
			&complaint.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		complaints = append(complaints, complaint)
	}

	return complaints, nil
}

// UpdateComplaintStatus updates user's complaint status in the database.
func (r *AdminRepository) UpdateComplaintStatus(id int, status string) error {

	query := `
	UPDATE complaints
	SET status = $1,
	    updated_at = NOW()
	WHERE id = $2;
	`

	// Run the updates status execution.
	result, err := database.DB.Exec(
		context.Background(),
		query,
		status,
		id,
	)

	if err != nil {
		return err
	}

	// Return not found if no rows were updated.
	if result.RowsAffected() == 0 {
		return utils.ErrComplaintNotFound
	}

	return nil
}

// GetAllUsers retrieves all registered user rows from the database.
func (r *AdminRepository) GetAllUsers() ([]dto.AdminUserResponse, error) {

	query := `
	SELECT
		id,
		name,
		email,
		role,
		is_active,
		created_at,
		updated_at
	FROM users
	ORDER BY id ASC;
	`

	// Run query to fetch all user accounts.
	rows, err := database.DB.Query(
		context.Background(),
		query,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []dto.AdminUserResponse

	// Iterate and scan results into user list.
	for rows.Next() {

		var user dto.AdminUserResponse

		// Scan database record fields.
		err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.Role,
			&user.IsActive,
			&user.CreatedAt,
			&user.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

// UpdateUserRole updates the role value of a user in the database.
func (r *AdminRepository) UpdateUserRole(id int, role string) error {
	query := `
		UPDATE users
		SET role = $1,
		    updated_at = NOW()
		WHERE id = $2;
	`

	// Run roles update execution.
	result, err := database.DB.Exec(
		context.Background(),
		query,
		role,
		id,
	)
	if err != nil {
		return err
	}

	// Return not found check.
	if result.RowsAffected() == 0 {
		return utils.ErrUserNotFound
	}

	return nil
}

// UpdateUserStatus updates the is_active flag of a user in the database.
func (r *AdminRepository) UpdateUserStatus(id int, isActive bool) error {
	query := `
		UPDATE users
		SET is_active = $1,
		    updated_at = NOW()
		WHERE id = $2;
	`

	// Run status modification statement.
	result, err := database.DB.Exec(
		context.Background(),
		query,
		isActive,
		id,
	)
	if err != nil {
		return err
	}

	// Return error if no records were altered.
	if result.RowsAffected() == 0 {
		return utils.ErrUserNotFound
	}

	return nil
}

// DeleteUser removes a user record from the database.
func (r *AdminRepository) DeleteUser(id int) error {

	query := `
	DELETE FROM users
	WHERE id = $1;
	`

	// Run removal statement.
	result, err := database.DB.Exec(
		context.Background(),
		query,
		id,
	)

	if err != nil {
		return err
	}

	// Return not found error if check counts zero.
	if result.RowsAffected() == 0 {
		return utils.ErrUserNotFound
	}

	return nil
}

// GetComplaintByID retrieves a single complaint from the database using its ID.
func (r *AdminRepository) GetComplaintByID(id int) (*models.Complaint, error) {

	complaint := &models.Complaint{}

	query := `
	SELECT
		id,
		user_id,
		title,
		description,
		status,
		created_at,
		updated_at
	FROM complaints
	WHERE id = $1;
	`

	// Run query to fetch the complaint details.
	err := database.DB.QueryRow(
		context.Background(),
		query,
		id,
	).Scan(
		&complaint.ID,
		&complaint.UserID,
		&complaint.Title,
		&complaint.Description,
		&complaint.Status,
		&complaint.CreatedAt,
		&complaint.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return complaint, nil
}

// GetComplaintEmailData retrieves the user email, username, and complaint details needed for email notifications.
func (r *AdminRepository) GetComplaintEmailData(id int) (*dto.ComplaintEmailData, error) {

	query := `
	SELECT
		u.name,
		u.email,
		c.title,
		c.status
	FROM complaints c
	INNER JOIN users u
		ON c.user_id = u.id
	WHERE c.id = $1;
	`

	data := &dto.ComplaintEmailData{}

	// Run query to fetch email notification details.
	err := database.DB.QueryRow(
		context.Background(),
		query,
		id,
	).Scan(
		&data.UserName,
		&data.UserEmail,
		&data.ComplaintTitle,
		&data.Status,
	)

	if err != nil {
		return nil, err
	}

	return data, nil
}
