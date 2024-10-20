package config

import (
	"encoding/json"
	"log"
	"os"

	"GO-API/models"
)

// SeedData seeds the database with initial data
func SeedData() {
	// Seed albums from JSON file
	if err := SeedAlbumsFromFile("albums.json"); err != nil {
		log.Printf("Error seeding albums: %v", err)
	}

	// Seed reviews from JSON file
	if err := SeedReviewsFromFile("reviews.json"); err != nil {
		log.Printf("Error seeding reviews: %v", err)
	}
}

func SeedAlbumsFromFile(filename string) error {
	file, err := os.Open(filename) // Open the file
	if err != nil {
		return err
	}
	defer file.Close() // Ensure the file is closed after reading

	var albums []models.Album
	if err := json.NewDecoder(file).Decode(&albums); err != nil { // Use json.Decoder
		return err
	}

	for _, album := range albums {
		DB.FirstOrCreate(&album, models.Album{Title: album.Title, Artist: album.Artist})
	}

	return nil
}

func SeedReviewsFromFile(filename string) error {
	file, err := os.Open(filename) // Open the file
	if err != nil {
		return err
	}
	defer file.Close() // Ensure the file is closed after reading

	var reviews []models.Review
	if err := json.NewDecoder(file).Decode(&reviews); err != nil { // Use json.Decoder
		return err
	}

	for _, review := range reviews {
		DB.FirstOrCreate(&review, models.Review{AlbumID: review.AlbumID, Content: review.Content})
	}

	return nil
}
