package utils

import "testing"

// TestHashPassword verifies that a plaintext password is
// successfully hashed and differs from the original value.
func TestHashPassword(t *testing.T) {
	password := "123456"

	hashedPassword, err := HashPassword(password)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if hashedPassword == "" {
		t.Fatal("expected hashed password, got empty string")
	}

	if hashedPassword == password {
		t.Fatal("password should be hashed, but it matches the original password")
	}
}

// TestCheckPassword verifies that the original password
// matches its hashed representation.
func TestCheckPassword(t *testing.T) {
	password := "123456"

	hashedPassword, _ := HashPassword(password)

	if !CheckPassword(password, hashedPassword) {
		t.Fatal("expected password to match")
	}
}

// TestCheckPasswordWrongPassword verifies that password
// verification fails when an incorrect password is provided.
func TestCheckPasswordWrongPassword(t *testing.T) {
	password := "123456"

	hashedPassword, _ := HashPassword(password)

	if CheckPassword("wrongpassword", hashedPassword) {
		t.Fatal("expected password check to fail")
	}
}