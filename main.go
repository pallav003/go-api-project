package main

import (
	"go-project/models"
	"go-project/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Connect to database
	models.ConnectDatabase()
	print("database connected")

	// Set up routes
	routes.UserRoutes(r)

	// Start server
	r.Run(":8080")
}
