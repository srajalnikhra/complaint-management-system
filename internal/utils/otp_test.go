package utils

import (
	"regexp"
	"testing"
)

// TestGenerateOTP verifies that the generated OTP
// is six digits long and contains only numeric characters.
func TestGenerateOTP(t *testing.T) {
	otp := GenerateOTP()

	if len(otp) != 6 {
		t.Errorf("expected OTP length 6, got %d", len(otp))
	}

	matched, err := regexp.MatchString(`^\d{6}$`, otp)
	if err != nil {
		t.Fatalf("failed to validate OTP format: %v", err)
	}

	if !matched {
		t.Errorf("expected numeric OTP, got %s", otp)
	}
}

// TestGenerateOTPUnique verifies that two consecutively
// generated OTPs are different.
func TestGenerateOTPUnique(t *testing.T) {
	otp1 := GenerateOTP()
	otp2 := GenerateOTP()

	if otp1 == otp2 {
		t.Errorf("expected different OTPs, both were %s", otp1)
	}
}