package database

import "github.com/srajalnikhra/complaint-management-system/internal/config"

// Initialize connects to the database, runs migrations, and seeds the default admin user.
func Initialize(dbConfig config.DBConfig) {
	// Connect to the database.
	ConnectDB(dbConfig)

	// Run database migrations.
	RunMigrations()

	// Create the default administrator account.
	BootstrapAdmin()
}
