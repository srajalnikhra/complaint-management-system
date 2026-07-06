package repositories

import (
	"context"

	"github.com/srajalnikhra/complaint-management-system/internal/database"
	"github.com/srajalnikhra/complaint-management-system/internal/models"
)

type AdminRepository struct{}

func NewAdminRepository() *AdminRepository {
	return &AdminRepository{}
}

func (r *AdminRepository) GetAllComplaints(page, limit int, search string) ([]models.Complaint, error) {

	offset := (page - 1) * limit

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
	WHERE
		title ILIKE '%' || $1 || '%'
		OR description ILIKE '%' || $1 || '%'
	ORDER BY id DESC
	LIMIT $2 OFFSET $3;
	`

	rows, err := database.DB.Query(
		context.Background(),
		query,
		search,
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

	_, err := database.DB.Exec(
		context.Background(),
		query,
		status,
		id,
	)

	return err
}