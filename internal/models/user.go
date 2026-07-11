package models

import "time"

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
