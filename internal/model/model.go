package model

import (
	"time"
)

// URL is the model for the Url table
type URL struct {
	ID        uint32 `gorm:"primary_key"`
	CreatedAt time.Time
	ShortURL  string `gorm:"index"`
	LongURL   string
}