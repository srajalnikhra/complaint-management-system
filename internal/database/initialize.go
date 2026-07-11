package database

import "github.com/srajalnikhra/complaint-management-system/internal/config"

func Initialize(dbConfig config.DBConfig) {
	ConnectDB(dbConfig)
	RunMigrations()
	BootstrapAdmin()
}
