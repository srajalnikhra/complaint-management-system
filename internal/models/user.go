package models

import "time"

// User maps to the users database record, representing customers or administrator accounts.
type User struct {
	ID       int
	Name     string
	Email    string
	Password string
	Role     string
	IsActive bool

	OTPHash      string
	OTPExpiresAt *time.Time

	CreatedAt time.Time
	UpdatedAt time.Time

	Token string `json:"token,omitempty"`
}
