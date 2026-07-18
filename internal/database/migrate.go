package database

import (
	"context"
	"log"
)

// RunMigrations sets up the database schema by creating the necessary tables and columns if they do not exist.
func RunMigrations() {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		email VARCHAR(255) UNIQUE NOT NULL,
		password TEXT NOT NULL,
		role VARCHAR(20) NOT NULL DEFAULT 'user',
		is_active BOOLEAN DEFAULT TRUE,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	`

	// Create the users table.
	_, err := DB.Exec(context.Background(), query)
	if err != nil {
		log.Fatal("Failed to run migrations:", err)
	}

	alterUsers := `
ALTER TABLE users
ADD COLUMN IF NOT EXISTS otp_hash TEXT,
ADD COLUMN IF NOT EXISTS otp_expires_at TIMESTAMPTZ;
`

	// Add OTP fields to the users table for the password reset flow.
	_, err = DB.Exec(context.Background(), alterUsers)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Database migrations completed successfully")

	complaintsTable := `
CREATE TABLE IF NOT EXISTS complaints(
	id SERIAL PRIMARY KEY,
	user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
	title VARCHAR(255) NOT NULL,
	description TEXT NOT NULL,
	status VARCHAR(20) NOT NULL DEFAULT 'PENDING',
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);`

	// Create the complaints table.
	_, err = DB.Exec(context.Background(), complaintsTable)
	if err != nil {
		log.Fatal(err)
	}
}
