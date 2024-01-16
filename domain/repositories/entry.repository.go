package repositories

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"go-graphql-cli/domain/models/entities"
)

type (
    EntryRepository interface {
        GetEntries() ([]entities.Entry, error)
    }
    entryRepositoryImpl struct{}
)

func NewEntryRepository() EntryRepository {
    return &entryRepositoryImpl{}
}

func (r *entryRepositoryImpl) GetEntries() ([]entities.Entry, error) {
    err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("Error loading .env file: %w", err)
	}

	postgresUser := os.Getenv("POSTGRES_USER")
    postgresPassword := os.Getenv("POSTGRES_PASSWORD")
    postgresDb := os.Getenv("POSTGRES_DB")
    postgresHost := os.Getenv("POSTGRES_HOST")
    postgresPort := os.Getenv("POSTGRES_PORT")

    dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", postgresUser, postgresPassword, postgresHost, postgresPort, postgresDb)
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        return nil, fmt.Errorf("Failed to connect to database: %w", err)
    }

    var entries []entities.Entry
    if err := db.Find(&entries).Error; err != nil {
        return nil, fmt.Errorf("Failed to get entries: %w", err)
    }

    return entries, nil
}
