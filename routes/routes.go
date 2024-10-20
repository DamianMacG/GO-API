// routes/routes.go
package routes

import (
	"GO-API/controllers"
	"github.com/gin-gonic/gin"
)

// RegisterRoutes sets up the API endpoints
func RegisterRoutes(router *gin.Engine) {
	router.GET("/albums", controllers.GetAlbums)
	router.GET("/albums/:id", controllers.GetAlbumByID)
	router.POST("/albums", controllers.PostAlbums)
	router.PUT("/albums/:id", controllers.UpdateAlbum)
	router.DELETE("/albums/:id", controllers.DeleteAlbum)

	// Reviews routes
	router.GET("/albums/:album_id/reviews", controllers.GetReviews) // Get reviews for an album
	router.POST("/albums/:album_id/reviews", controllers.PostReview) // Add a review for an album
}

