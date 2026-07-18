// Package dto houses Data Transfer Objects containing parsing tags for incoming and outgoing HTTP payloads.
package dto

// LoginRequest defines credentials variables submitted during user login.
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginResponse formats the success response for authentication.
type LoginResponse struct {
	User  UserResponse `json:"user"`
	Token string       `json:"token"`
}
