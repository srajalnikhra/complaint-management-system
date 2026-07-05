package dto

type UpdateComplaintRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}