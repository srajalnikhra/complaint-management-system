package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// LoadEnv reads environment variables from the .env file in development.
// If the file is missing, it falls back to the system environment variables.
func LoadEnv() {
	// Try to load the .env file. If it fails, default to system env.
	if err := godotenv.Load(".env"); err != nil {
		log.Println(".env file not found, using environment variables")
	}
}

// LoadAppConfig reads application and email configurations from the environment.
func LoadAppConfig() AppConfig {
	// Load general settings and mail server credentials.
	return AppConfig{
		Name: os.Getenv("APP_NAME"),
		Env:  os.Getenv("APP_ENV"),
		Port: os.Getenv("APP_PORT"),

		SMTPHost:     os.Getenv("SMTP_HOST"),
		SMTPPort:     os.Getenv("SMTP_PORT"),
		SMTPEmail:    os.Getenv("SMTP_EMAIL"),
		SMTPPassword: os.Getenv("SMTP_PASSWORD"),
	}
}

// LoadDBConfig reads database host details and credentials from the environment.
func LoadDBConfig() DBConfig {
	// Load connection settings for PostgreSQL.
	return DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Name:     os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}
}

// LoadJWTConfig reads the secret key and token lifetime for signing JWTs.
func LoadJWTConfig() JWTConfig {
	// Load JWT configuration settings.
	return JWTConfig{
		Secret: os.Getenv("JWT_SECRET"),
		Expiry: os.Getenv("JWT_EXPIRY"),
	}
}
