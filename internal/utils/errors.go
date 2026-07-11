package utils

import "errors"

var (
	ErrUserAlreadyExists = errors.New("email already exists")

	ErrInvalidCredentials = errors.New("invalid email or password")

	ErrComplaintNotFound = errors.New("complaint not found")

	ErrUserNotFound = errors.New("user not found")
)
