package db

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"go-graphql-cli/domain/models/entities"
)

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
    err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("Error loading .env file: %w", err)
	}

	postgresUser := os.Getenv("POSTGRES_USER")
    postgresPassword := os.Getenv("POSTGRES_PASSWORD")
    postgresDb := os.Getenv("POSTGRES_DB")
    postgresHost := os.Getenv("POSTGRES_HOST")
    postgresPort := os.Getenv("POSTGRES_PORT")

    // connect to db
    dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", postgresUser, postgresPassword, postgresHost, postgresPort, postgresDb)
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        return nil, fmt.Errorf("Failed to connect to database: %w", err)
    }

    // migrate
    db.AutoMigrate(
        &entities.Entry{},
    )

	return db, nil
}
