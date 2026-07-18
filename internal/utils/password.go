package utils

import "golang.org/x/crypto/bcrypt"

// HashPassword securely hashes plain password strings using the bcrypt algorithm.
func HashPassword(password string) (string, error) {
	// Encrypt raw password string using standard bcrypt hash algorithm and default work factor.
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPassword verifies that a plaintext password candidate matches the hashed version.
func CheckPassword(password, hashedPassword string) bool {
	// Compare stored encrypted value against potential plaintext candidate.
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
