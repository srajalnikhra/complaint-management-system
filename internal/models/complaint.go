// Package models contains the structural definitions mapping to database schemas.
package models

import "time"

// Complaint holds the data schema mapping to the postgres complaints table.
type Complaint struct {
	ID          int
	UserID      int
	Title       string
	Description string
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
