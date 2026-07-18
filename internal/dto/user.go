package dto

// RegisterUserRequest represents the payload payload for signing up.
type RegisterUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
