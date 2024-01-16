package entities

import "time"

type Entry struct {
	ID        string    `gorm:"column:id;primary_key"`
	Name      string    `gorm:"column:name"`
	CreatedAt time.Time `gorm:"column:created_at"`
}
