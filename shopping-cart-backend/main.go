package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"shopping-cart-backend/controllers"
	"shopping-cart-backend/database"
	"shopping-cart-backend/middleware"
)

func main() {
	// Connect to the database and run migrations
	database.ConnectDatabase()

	// Initialize Gin router
	r := gin.Default()

	// Add CORS middleware (crucial for frontend to communicate with backend)
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*") // Allow requests from any origin (for development)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// Public routes (no authentication required)
	// These routes correctly use 'r.' because they are not part of any group
	r.POST("/users", controllers.CreateUser)
	r.POST("/users/login", controllers.LoginUser)

	// Authenticated routes (require a valid JWT token)
	// These routes will pass through the AuthMiddleware for token validation.
	authorized := r.Group("/")                  // Create the group
	authorized.Use(middleware.AuthMiddleware()) // Apply middleware to the group
	{
		// IMPORTANT: All routes defined within this block MUST use 'authorized.'
		// This is the core fix for the "cannot slice" error.

		// User routes
		authorized.GET("/users", controllers.GetUsers) // Corrected: Uses 'authorized.'

		// Item routes
		authorized.POST("/items", controllers.CreateItem) // Corrected: Uses 'authorized.'
		authorized.GET("/items", controllers.GetItems)    // Corrected: Uses 'authorized.'

		// Cart routes
		authorized.POST("/carts", controllers.AddItemsToCart) // Corrected: Uses 'authorized.'
		authorized.GET("/carts", controllers.GetCartByUserID) // Corrected: Uses 'authorized.'

		// Order routes
		authorized.POST("/orders", controllers.CreateOrder)      // Corrected: Uses 'authorized.'
		authorized.GET("/orders", controllers.GetOrdersByUserID) // Corrected: Uses 'authorized.'
	}

	log.Println("Server starting on :8080")
	log.Fatal(r.Run(":8080")) // Run the server on port 8080
}
