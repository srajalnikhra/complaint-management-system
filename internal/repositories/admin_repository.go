package repositories

import (
	"context"
	"fmt"

	"github.com/srajalnikhra/complaint-management-system/internal/database"
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
