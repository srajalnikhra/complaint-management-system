package database

import (
	"context"
	"log"

	"github.com/srajalnikhra/complaint-management-system/internal/models"
	"github.com/srajalnikhra/complaint-management-system/internal/utils"
)

// BootstrapAdmin creates a default admin account on startup if no admin users exist yet.
func BootstrapAdmin() {
	var count int

	// Check if there are any existing admin users in the database.
	err := DB.QueryRow(
		context.Background(),
		`SELECT COUNT(*) FROM users WHERE role = $1`,
		models.RoleAdmin,
	).Scan(&count)

	if err != nil {
		log.Fatal("Failed to check admin:", err)
	}

	if count > 0 {
		log.Println("Admin already exists")
		return
	}

	// Hash the default password before saving it.
	hashedPassword, err := utils.HashPassword("123456")
	if err != nil {
		log.Fatal("Failed to hash admin password:", err)
	}

	// Insert the default admin record.
	_, err = DB.Exec(
		context.Background(),
		`
		INSERT INTO users
		(name, email, password, role, is_active)
		VALUES ($1, $2, $3, $4, $5)
		`,
		"Administrator",
		"admin@gmail.com",
		hashedPassword,
		models.RoleAdmin,
		true,
	)

	if err != nil {
		log.Fatal("Failed to create admin:", err)
	}

	log.Println("Default admin created")
}
