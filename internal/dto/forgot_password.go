package dto

// ForgotPasswordRequest represents the email submission for password recovery.
type ForgotPasswordRequest struct {
	Email string `json:"email"`
}

// VerifyOTPRequest contains the OTP verification payload.
type VerifyOTPRequest struct {
	Email string `json:"email"`
	OTP   string `json:"otp"`
}

// ResetPasswordRequest carries the user's new password and OTP code.
type ResetPasswordRequest struct {
	Email       string `json:"email"`
	OTP         string `json:"otp"`
	NewPassword string `json:"new_password"`
}
