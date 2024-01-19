package entities

import (
	"time"

	"gorm.io/gorm"
)

type Entry struct {
	gorm.Model
	ID        string    `gorm:"column:id;primary_key"`
	Name      string    `gorm:"column:name"`
	CreatedAt time.Time `gorm:"column:created_at"`
}
