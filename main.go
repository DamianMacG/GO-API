package main

import (
	"GO-API/config"
	"GO-API/routes"
	"log"

	"github.com/joho/godotenv"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect to the database
	config.ConnectDatabase() 
	config.SeedData() // Seed the database with initial data

	// Set up the router
	router := gin.Default()
	routes.RegisterRoutes(router) // Register your routes

	router.Run("localhost:8080") // Run the server
}
