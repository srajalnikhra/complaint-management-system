// Package dto houses Data Transfer Objects containing parsing tags for incoming and outgoing HTTP payloads.
package dto

import "time"

// CreateComplaintRequest models the request body used for submitting a new complaint.
type CreateComplaintRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

// ComplaintResponse formats complaint details for API responses.
type ComplaintResponse struct {
	ID          int
	Title       string
	Description string
	Status      string
	CreatedAt   time.Time
}

// PaginationQuery handles query parameters for listing records.
type PaginationQuery struct {
	Page  int
	Limit int
}
