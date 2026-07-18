package models

// Role definitions for authorization scopes.
const (
	// RoleUser represents regular application users who submit complaints.
	RoleUser = "user"
	// RoleAdmin represents system administrators with access to manage users and complaints.
	RoleAdmin = "admin"
)
