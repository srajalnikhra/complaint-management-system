package dto

// UpdateComplaintRequest represents the payload for user edits to a complaint.
type UpdateComplaintRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

// UpdateComplaintStatusRequest handles the status update provided by the admin.
type UpdateComplaintStatusRequest struct {
	Status string `json:"status"`
}
