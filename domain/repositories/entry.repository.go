package repositories

import (
	"fmt"

	"gorm.io/gorm"

	"go-graphql-cli/domain/models/entities"
)

type (
    EntryRepository interface {
        GetEntries() ([]entities.Entry, error)
        GetEntry(id string) (*entities.Entry, error)
        CreateEntry(entry *entities.Entry) error
    }
    entryRepositoryImpl struct{
        db *gorm.DB
    }
)

func NewEntryRepository(db *gorm.DB) EntryRepository {
    return &entryRepositoryImpl{
        db: db,
    }
}

func (r *entryRepositoryImpl) GetEntries() ([]entities.Entry, error) {
    var entries []entities.Entry
    if err := r.db.Find(&entries).Error; err != nil {
        return nil, fmt.Errorf("Failed to get entries: %w", err)
    }

    return entries, nil
}

func (r *entryRepositoryImpl) GetEntry(id string) (*entities.Entry, error) {
    var entry entities.Entry
    if err := r.db.Where("id = ?", id).First(&entry).Error; err != nil {
        return nil, fmt.Errorf("Failed to get entry: %w", err)
    }

    return &entry, nil
}

func (r *entryRepositoryImpl) CreateEntry(entry *entities.Entry) error {
    if err := r.db.Create(&entry).Error; err != nil {
        return fmt.Errorf("Failed to create entry: %w", err)
    }

    return nil
}
