package repositories

import (
	"context"

	"github.com/srajalnikhra/complaint-management-system/internal/database"
	"github.com/srajalnikhra/complaint-management-system/internal/models"
	"github.com/srajalnikhra/complaint-management-system/internal/utils"
)

// ComplaintRepository manages direct database operations for user complaints.
type ComplaintRepository struct{}

// NewComplaintRepository creates a new instance of ComplaintRepository.
func NewComplaintRepository() *ComplaintRepository {
	return &ComplaintRepository{}
}

// Create saves a new complaint record in the database, populating status and timestamps.
func (r *ComplaintRepository) Create(c *models.Complaint) error {

	query := `
	INSERT INTO complaints(user_id,title,description)
	VALUES($1,$2,$3)
	RETURNING id,status,created_at,updated_at;
	`

	// Run the insert query and scan returned columns.
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

// GetByUserID retrieves all complaints submitted by a given user.
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

	// Run query to fetch user complaints.
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

		// Scan row details into complaint struct.
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

// GetByID retrieves a single complaint by its unique ID.
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

	// Run query to fetch specific complaint details.
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
		return nil, utils.ErrComplaintNotFound
	}

	return &complaint, nil
}

// Update modifies an existing complaint's title and description, returning the updated timestamp.
func (r *ComplaintRepository) Update(c *models.Complaint) error {
	query := `
	UPDATE complaints
	SET title = $1,
		description = $2,
		updated_at = CURRENT_TIMESTAMP
	WHERE id = $3
	RETURNING updated_at;
	`

	// Run update statement and scan the new updated_at field.
	return database.DB.QueryRow(
		context.Background(),
		query,
		c.Title,
		c.Description,
		c.ID,
	).Scan(&c.UpdatedAt)
}

// Delete removes a complaint record from the database.
func (r *ComplaintRepository) Delete(id int) error {
	query := `
	DELETE FROM complaints
	WHERE id = $1;
	`

	// Run the delete query.
	_, err := database.DB.Exec(
		context.Background(),
		query,
		id,
	)

	return err
}
