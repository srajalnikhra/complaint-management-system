package database

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/srajalnikhra/complaint-management-system/internal/config"
)

// ConnectDB connects to PostgreSQL and assigns the connection pool to the global DB variable.
func ConnectDB(dbConfig config.DBConfig) {
	// Format the connection string (DSN).
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Name,
		dbConfig.SSLMode,
	)

	// Connect to the database and initialize the pool.
	db, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatal("Failed to create database pool:", err)
	}

	// Ping the database to verify the connection works.
	if err := db.Ping(context.Background()); err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Store the connection pool in the global variable.
	DB = db

	log.Println("Database connected successfully")
}
