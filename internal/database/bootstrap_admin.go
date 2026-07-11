package database

import (
	"context"
	"log"

	"github.com/srajalnikhra/complaint-management-system/internal/models"
	"github.com/srajalnikhra/complaint-management-system/internal/utils"
)

func BootstrapAdmin() {
	var count int

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

	hashedPassword, err := utils.HashPassword("123456")
	if err != nil {
		log.Fatal("Failed to hash admin password:", err)
	}

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
