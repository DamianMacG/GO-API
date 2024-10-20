// controllers/review.go
package controllers

import (
	"net/http"
	"GO-API/config"
	"GO-API/models"
	"github.com/gin-gonic/gin"
)

// GetReviews retrieves all reviews for a specific album
func GetReviews(c *gin.Context) {
	albumID := c.Param("album_id")
	var reviews []models.Review

	if err := config.DB.Where("album_id = ?", albumID).Find(&reviews).Error; err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "no reviews found"})
		return
	}

	c.IndentedJSON(http.StatusOK, reviews)
}

// PostReview adds a new review for an album
func PostReview(c *gin.Context) {
	var newReview models.Review

	if err := c.BindJSON(&newReview); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		return
	}

	config.DB.Create(&newReview)
	c.IndentedJSON(http.StatusCreated, newReview)
}
