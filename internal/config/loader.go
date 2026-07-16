package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	if err := godotenv.Load(".env"); err != nil {
		log.Println(".env file not found, using environment variables")
	}
}

func LoadAppConfig() AppConfig {
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

func LoadDBConfig() DBConfig {
	return DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Name:     os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}
}

func LoadJWTConfig() JWTConfig {
	return JWTConfig{
		Secret: os.Getenv("JWT_SECRET"),
		Expiry: os.Getenv("JWT_EXPIRY"),
	}
}
