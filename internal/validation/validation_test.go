package validation

import (
	"testing"

	"github.com/srajalnikhra/complaint-management-system/internal/models"
)

// TestValidateRegister verifies registration input validation.
func TestValidateRegister(t *testing.T) {
	if msg := ValidateRegister("Srajal", "srajal@gmail.com", "123456"); msg != "" {
		t.Errorf("expected no validation error, got %s", msg)
	}

	if msg := ValidateRegister("", "srajal@gmail.com", "123456"); msg != "Name is required" {
		t.Errorf("expected 'Name is required', got %s", msg)
	}
}

// TestValidateLogin verifies login input validation.
func TestValidateLogin(t *testing.T) {
	if msg := ValidateLogin("srajal@gmail.com", "123456"); msg != "" {
		t.Errorf("expected no validation error, got %s", msg)
	}

	if msg := ValidateLogin("invalid-email", "123456"); msg != "Invalid email address" {
		t.Errorf("expected invalid email error, got %s", msg)
	}
}

// TestValidateComplaint verifies complaint field validation.
func TestValidateComplaint(t *testing.T) {
	if msg := ValidateComplaint("Internet Issue", "WiFi is not working"); msg != "" {
		t.Errorf("expected no validation error, got %s", msg)
	}

	if msg := ValidateComplaint("", "WiFi is not working"); msg != "Title is required" {
		t.Errorf("expected title required error, got %s", msg)
	}
}

// TestValidateComplaintStatus verifies complaint status validation.
func TestValidateComplaintStatus(t *testing.T) {
	if msg := ValidateComplaintStatus("PENDING"); msg != "" {
		t.Errorf("expected no validation error, got %s", msg)
	}

	if msg := ValidateComplaintStatus("DONE"); msg != "Invalid complaint status" {
		t.Errorf("expected invalid status error, got %s", msg)
	}
}

// TestValidateUserRole verifies user role validation.
func TestValidateUserRole(t *testing.T) {
	if msg := ValidateUserRole(models.RoleUser); msg != "" {
		t.Errorf("expected no validation error, got %s", msg)
	}

	if msg := ValidateUserRole("manager"); msg != "role must be either user or admin" {
		t.Errorf("expected invalid role error, got %s", msg)
	}
}

// TestValidateForgotPassword verifies forgot password validation.
func TestValidateForgotPassword(t *testing.T) {
	if msg := ValidateForgotPassword("srajal@gmail.com"); msg != "" {
		t.Errorf("expected no validation error, got %s", msg)
	}

	if msg := ValidateForgotPassword("invalid-email"); msg != "Invalid email address" {
		t.Errorf("expected invalid email error, got %s", msg)
	}
}

// TestValidateVerifyOTP verifies OTP validation.
func TestValidateVerifyOTP(t *testing.T) {
	if msg := ValidateVerifyOTP("srajal@gmail.com", "123456"); msg != "" {
		t.Errorf("expected no validation error, got %s", msg)
	}

	if msg := ValidateVerifyOTP("srajal@gmail.com", "123"); msg != "OTP must be 6 digits" {
		t.Errorf("expected OTP length error, got %s", msg)
	}
}

// TestValidateResetPassword verifies password reset validation.
func TestValidateResetPassword(t *testing.T) {
	if msg := ValidateResetPassword("srajal@gmail.com", "123456", "abcdef"); msg != "" {
		t.Errorf("expected no validation error, got %s", msg)
	}

	if msg := ValidateResetPassword("srajal@gmail.com", "123456", "123"); msg != "Password must be at least 6 characters long" {
		t.Errorf("expected password length error, got %s", msg)
	}
}