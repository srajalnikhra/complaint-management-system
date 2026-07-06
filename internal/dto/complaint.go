package dto

import "time"

type CreateComplaintRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type ComplaintResponse struct {
	ID          int
	Title       string
	Description string
	Status      string
	CreatedAt   time.Time
}

type PaginationQuery struct {
	Page  int
	Limit int
}
