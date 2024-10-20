// config/seed.go
package config

import (
	"GO-API/models"
)

// SeedData seeds the database with initial data
func SeedData() {
	albums := []models.Album{
		{Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
		{Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	}

	for _, album := range albums {
		DB.FirstOrCreate(&album, models.Album{Title: album.Title, Artist: album.Artist})
	}

	// Seed reviews for the albums
	reviews := []models.Review{
		{AlbumID: 1, Content: "A masterpiece!", Rating: 5},
		{AlbumID: 1, Content: "Great for jazz lovers!", Rating: 4},
		{AlbumID: 2, Content: "Not my favorite.", Rating: 2},
	}

	for _, review := range reviews {
		// Check if the review already exists before creating
		DB.FirstOrCreate(&review, models.Review{AlbumID: review.AlbumID, Content: review.Content})
	}
}
