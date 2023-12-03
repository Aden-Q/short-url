package model

import (
	"time"
)

// Model is the common database model specification
type Model struct {
	ID        uint64 `gorm:"primary_key"`
	CreatedAt time.Time
}
