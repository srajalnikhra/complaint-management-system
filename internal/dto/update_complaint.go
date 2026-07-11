package dto

type UpdateComplaintRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type UpdateComplaintStatusRequest struct {
	Status string `json:"status"`
}
