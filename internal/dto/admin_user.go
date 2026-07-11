package dto

import "time"

type AdminUserResponse struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UpdateUserRoleRequest struct {
	Role string `json:"role"`
}

type UpdateUserStatusRequest struct {
	IsActive bool `json:"is_active"`
}
