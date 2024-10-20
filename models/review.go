// models/review.go
package models

import (
	"gorm.io/gorm"
)

// Review represents a review of an album
type Review struct {
	gorm.Model
	AlbumID uint   `json:"album_id"` // Foreign key referencing the Album
	Content  string `json:"content"`
	Rating   int    `json:"rating"` // Example rating field
}

