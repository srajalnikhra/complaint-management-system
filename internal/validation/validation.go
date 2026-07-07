package validation

import (
	"net/mail"
	"strings"
)

func ValidateRegister(name, email, password string) string {

	if strings.TrimSpace(name) == "" {
		return "Name is required"
	}

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

	if len(password) < 6 {
		return "Password must be at least 6 characters long"
	}

	return ""
}

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

func ValidateComplaint(title, description string) string {

	if strings.TrimSpace(title) == "" {
		return "Title is required"
	}

	if strings.TrimSpace(description) == "" {
		return "Description is required"
	}

	return ""
}

func ValidateComplaintStatus(status string) string {

	switch status {
	case "PENDING", "IN_PROGRESS", "RESOLVED":
		return ""
	default:
		return "Invalid complaint status"
	}
}