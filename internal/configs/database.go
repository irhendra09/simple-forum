package configs

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func env(key, fallback string) string {
	v := os.Getenv(key)
	if v == "" {
		return fallback
	}
	return v
}

// ConnectDB builds DSN from environment variables or AppConfig when loaded.
// Priority: environment variables -> AppConfig values -> sensible defaults.
func ConnectDB() *gorm.DB {
	// environment overrides
	host := env("DB_HOST", "")
	port := env("DB_PORT", "")
	user := env("DB_USER", "")
	password := env("DB_PASSWORD", "")
	dbname := env("DB_NAME", "")
	sslmode := env("DB_SSLMODE", "")

	// fallback to AppConfig when environment not provided
	if AppConfig != nil {
		if host == "" && AppConfig.Database.Host != "" {
			host = AppConfig.Database.Host
		}
		if port == "" && AppConfig.Database.Port != "" {
			port = AppConfig.Database.Port
		}
		if user == "" && AppConfig.Database.User != "" {
			user = AppConfig.Database.User
		}
		if password == "" && AppConfig.Database.Password != "" {
			password = AppConfig.Database.Password
		}
		if dbname == "" && AppConfig.Database.Name != "" {
			dbname = AppConfig.Database.Name
		}
		if sslmode == "" && AppConfig.Database.SSLMode != "" {
			sslmode = AppConfig.Database.SSLMode
		}
	}

	// final fallbacks
	if host == "" {
		host = "localhost"
	}
	if port == "" {
		port = "5432"
	}
	if user == "" {
		user = "postgres"
	}
	if password == "" {
		password = "admin123"
	}
	if dbname == "" {
		dbname = "forum_db"
	}
	if sslmode == "" {
		sslmode = "disable"
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbname, sslmode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	return db
}
