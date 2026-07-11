package repositories

import (
	"context"
	"fmt"

	"github.com/srajalnikhra/complaint-management-system/internal/database"
	"github.com/srajalnikhra/complaint-management-system/internal/dto"
	"github.com/srajalnikhra/complaint-management-system/internal/models"
	"github.com/srajalnikhra/complaint-management-system/internal/utils"
)

type AdminRepository struct{}

func NewAdminRepository() *AdminRepository {
	return &AdminRepository{}
}

func (r *AdminRepository) GetAllComplaints(page, limit int, search, status, sort, order string) ([]models.Complaint, error) {

	offset := (page - 1) * limit

	allowedSorts := map[string]bool{
		"id":         true,
		"created_at": true,
		"updated_at": true,
	}

	if !allowedSorts[sort] {
		sort = "id"
	}

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

	for rows.Next() {

		var complaint models.Complaint

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

func (r *AdminRepository) UpdateComplaintStatus(id int, status string) error {

	query := `
	UPDATE complaints
	SET status = $1,
	    updated_at = NOW()
	WHERE id = $2;
	`

	result, err := database.DB.Exec(
		context.Background(),
		query,
		status,
		id,
	)

	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return utils.ErrComplaintNotFound
	}

	return nil
}

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

	rows, err := database.DB.Query(
		context.Background(),
		query,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []dto.AdminUserResponse

	for rows.Next() {

		var user dto.AdminUserResponse

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

func (r *AdminRepository) UpdateUserRole(id int, role string) error {
	query := `
		UPDATE users
		SET role = $1,
		    updated_at = NOW()
		WHERE id = $2;
	`

	result, err := database.DB.Exec(
		context.Background(),
		query,
		role,
		id,
	)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return utils.ErrUserNotFound
	}

	return nil
}

func (r *AdminRepository) UpdateUserStatus(id int, isActive bool) error {
	query := `
		UPDATE users
		SET is_active = $1,
		    updated_at = NOW()
		WHERE id = $2;
	`

	result, err := database.DB.Exec(
		context.Background(),
		query,
		isActive,
		id,
	)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return utils.ErrUserNotFound
	}

	return nil
}

func (r *AdminRepository) DeleteUser(id int) error {

	query := `
	DELETE FROM users
	WHERE id = $1;
	`

	result, err := database.DB.Exec(
		context.Background(),
		query,
		id,
	)

	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return utils.ErrUserNotFound
	}

	return nil
}

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
