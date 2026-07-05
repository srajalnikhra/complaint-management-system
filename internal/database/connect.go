package database

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/srajalnikhra/complaint-management-system/internal/config"
)

func ConnectDB(dbConfig config.DBConfig) {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Name,
		dbConfig.SSLMode,
	)

	db, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatal("Failed to create database pool:", err)
	}

	if err := db.Ping(context.Background()); err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	DB = db

	log.Println("Database connected successfully")
}