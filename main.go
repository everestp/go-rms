package main

import (
	"os"

	"go-rms/database"
	"go-rms/middleware"
	"go-rms/routes"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)


var foodCollection *mongo.Collection = database.OpenCollections(database.Client ,"food")
func main() {
	// Connect to MongoDB


	// Optional: get a specific collection if needed globally
	// foodCollection := database.OpenCollections(database.Client, "food")

	// Set port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	// Create Gin router with default middleware
	router := gin.Default()

	// Register public routes first
	routes.UserRoutes(router)

	// Apply authentication middleware to protect the following routes
	router.Use(middleware.Authentication())

	// Register protected routes
	routes.FoodRouter(router)
	routes.TableRouter(router)
	routes.OrderRoutes(router)
	routes.InvoiceRoutes(router)

	// Start the server
	router.Run(":" + port)
}
