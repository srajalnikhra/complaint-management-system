// Package dto houses Data Transfer Objects containing parsing tags for incoming and outgoing HTTP payloads.
package dto

import "time"

// AdminUserResponse represents a sanitized user detail payload sent to administrators.
type AdminUserResponse struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// UpdateUserRoleRequest represents the payload for changing a user's role.
type UpdateUserRoleRequest struct {
	Role string `json:"role"`
}

// UpdateUserStatusRequest represents the payload for modifying a user's active status.
type UpdateUserStatusRequest struct {
	IsActive bool `json:"is_active"`
}
