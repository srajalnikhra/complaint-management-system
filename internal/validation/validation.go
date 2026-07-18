package validation

import (
	"net/mail"
	"strings"

	"github.com/srajalnikhra/complaint-management-system/internal/models"
)

// ValidateRegister asserts that the registration fields are present and correctly formatted.
func ValidateRegister(name, email, password string) string {

	if strings.TrimSpace(name) == "" {
		return "Name is required"
	}

	if strings.TrimSpace(email) == "" {
		return "Email is required"
	}

	// Validate email format using Go's built-in mail parser.
	_, err := mail.ParseAddress(email)
	if err != nil {
		return "Invalid email address"
	}

	if strings.TrimSpace(password) == "" {
		return "Password is required"
	}

	// Enforce a minimum password length pattern.
	if len(password) < 6 {
		return "Password must be at least 6 characters long"
	}

	return ""
}

// ValidateLogin checks credentials inputs.
func ValidateLogin(email, password string) string {

	if strings.TrimSpace(email) == "" {
		return "Email is required"
	}

	_, err := mail.ParseAddress(email)
	if err != nil {
		return "Invalid email address"
	}

	if strings.TrimSpace(password) == "" {
		return "Password is required"
	}

	return ""
}

// ValidateComplaint validates complaint fields.
func ValidateComplaint(title, description string) string {

	if strings.TrimSpace(title) == "" {
		return "Title is required"
	}

	if strings.TrimSpace(description) == "" {
		return "Description is required"
	}

	return ""
}

// ValidateComplaintStatus checks that status transitions correspond to known states.
func ValidateComplaintStatus(status string) string {

	switch status {
	case "PENDING", "IN_PROGRESS", "RESOLVED":
		return ""
	default:
		return "Invalid complaint status"
	}
}

// ValidateUserRole restricts role options.
func ValidateUserRole(role string) string {
	if role != models.RoleUser && role != models.RoleAdmin {
		return "role must be either user or admin"
	}

	return ""
}

// ValidateForgotPassword validates forgotten password request parameters.
func ValidateForgotPassword(email string) string {

	if strings.TrimSpace(email) == "" {
		return "Email is required"
	}

	_, err := mail.ParseAddress(email)
	if err != nil {
		return "Invalid email address"
	}

	return ""
}

// ValidateVerifyOTP validates verification parameters.
func ValidateVerifyOTP(email, otp string) string {

	if msg := ValidateForgotPassword(email); msg != "" {
		return msg
	}

	if strings.TrimSpace(otp) == "" {
		return "OTP is required"
	}

	if len(otp) != 6 {
		return "OTP must be 6 digits"
	}

	return ""
}

// ValidateResetPassword checks password reset inputs.
func ValidateResetPassword(email, otp, password string) string {

	if msg := ValidateVerifyOTP(email, otp); msg != "" {
		return msg
	}

	if strings.TrimSpace(password) == "" {
		return "Password is required"
	}

	if len(password) < 6 {
		return "Password must be at least 6 characters long"
	}

	return ""
}
