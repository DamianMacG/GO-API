// controllers/album.go
package controllers

import (
	"net/http"
	"GO-API/config"
	"GO-API/models"
	"github.com/gin-gonic/gin"
)

// GetAlbums responds with the list of all albums as JSON.
func GetAlbums(c *gin.Context) {
	var albums []models.Album
	config.DB.Find(&albums)
	c.IndentedJSON(http.StatusOK, albums)
}

// GetAlbumByID locates the album whose ID value matches the id parameter sent by the client, then returns that album as a response.
func GetAlbumByID(c *gin.Context) {
	id := c.Param("id")
	var album models.Album

	if err := config.DB.First(&album, id).Error; err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, album)
}

// PostAlbums adds an album from JSON received in the request body.
func PostAlbums(c *gin.Context) {
	var newAlbum models.Album

	if err := c.BindJSON(&newAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		return
	}

	config.DB.Create(&newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// UpdateAlbum updates an album
func UpdateAlbum(c *gin.Context) {
	id := c.Param("id")
	var updatedAlbum models.Album

	if err := c.BindJSON(&updatedAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		return
	}

	var album models.Album
	if err := config.DB.First(&album, id).Error; err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album not found"})
		return
	}

	// Update fields
	album.Title = updatedAlbum.Title
	album.Artist = updatedAlbum.Artist
	album.Price = updatedAlbum.Price

	config.DB.Save(&album)
	c.IndentedJSON(http.StatusOK, album)
}

// DeleteAlbum deletes an album
func DeleteAlbum(c *gin.Context) {
	id := c.Param("id")
	if err := config.DB.Delete(&models.Album{}, id).Error; err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Album deleted"})
}
