// Package utils houses helper utility libraries (security, formatters, and assertions).
package utils

import "errors"

// Standard domain-level error definitions to prevent hardcoded error messages throughout services and repositories.
var (
	// ErrUserAlreadyExists is thrown when someone registers an email that already exists.
	ErrUserAlreadyExists = errors.New("email already exists")

	// ErrInvalidCredentials indicates the email or password was wrong during login.
	ErrInvalidCredentials = errors.New("invalid email or password")

	// ErrComplaintNotFound means the assigned complaint ID was not located.
	ErrComplaintNotFound = errors.New("complaint not found")

	// ErrUserNotFound indicates the required user entry could not be found.
	ErrUserNotFound = errors.New("user not found")
)
