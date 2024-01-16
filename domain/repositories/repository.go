package repositories

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewRepository() {
	dsn := "postgres://postgres:@postgres:5432/go-graphql-cli?sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }
    log.Println("Connection Opened to Database: %v", db)
}
