package repositories

import (
	"context"

	"github.com/srajalnikhra/complaint-management-system/internal/database"
	"github.com/srajalnikhra/complaint-management-system/internal/models"
)

type ComplaintRepository struct{}

func NewComplaintRepository() *ComplaintRepository {
	return &ComplaintRepository{}
}

func (r *ComplaintRepository) Create(c *models.Complaint) error {

	query := `
	INSERT INTO complaints(user_id,title,description)
	VALUES($1,$2,$3)
	RETURNING id,status,created_at,updated_at;
	`

	return database.DB.QueryRow(
		context.Background(),
		query,
		c.UserID,
		c.Title,
		c.Description,
	).Scan(
		&c.ID,
		&c.Status,
		&c.CreatedAt,
		&c.UpdatedAt,
	)
}

func (r *ComplaintRepository) GetByUserID(userID int) ([]models.Complaint, error) {

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
	WHERE user_id = $1
	ORDER BY id DESC;
	`

	rows, err := database.DB.Query(
		context.Background(),
		query,
		userID,
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

func (r *ComplaintRepository) GetByID(id int) (*models.Complaint, error) {

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

	var complaint models.Complaint

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

	return &complaint, nil
}

func (r *ComplaintRepository) Update(c *models.Complaint) error {
	query := `
	UPDATE complaints
	SET title = $1,
		description = $2,
		updated_at = CURRENT_TIMESTAMP
	WHERE id = $3
	RETURNING updated_at;
	`

	return database.DB.QueryRow(
		context.Background(),
		query,
		c.Title,
		c.Description,
		c.ID,
	).Scan(&c.UpdatedAt)
}

func (r *ComplaintRepository) Delete(id int) error {
	query := `
	DELETE FROM complaints
	WHERE id = $1;
	`

	_, err := database.DB.Exec(
		context.Background(),
		query,
		id,
	)

	return err
}