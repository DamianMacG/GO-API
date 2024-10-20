package main

import (
	"GO-API/routes"
)

func main() {
	router := routes.SetupRouter()
	router.Run("localhost:8080")
}