package middleware

import (
	"fmt" // Used for fmt.Errorf
	"net/http"
	"strings"

	"shopping-cart-backend/utils" // Adjust path based on your module name

	"github.com/gin-gonic/gin"
)

// AuthMiddleware is a Gin middleware to authenticate requests using JWT tokens.
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the Authorization header from the request
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort() // Abort the request and do not proceed to handlers
			return
		}

		// Expect the token in "Bearer <token>" format
		if !strings.HasPrefix(tokenString, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format (expected Bearer token)"})
			c.Abort()
			return
		}
		tokenString = strings.TrimPrefix(tokenString, "Bearer ") // Extract the actual token string

		// Validate the token
		userID, err := utils.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("Invalid or expired token: %v", err)}) // Using fmt.Sprintf here
			c.Abort()
			return
		}

		// If valid, store the userID in the Gin context.
		// This makes the userID accessible in subsequent controller functions.
		c.Set("userID", userID)
		c.Next() // Proceed to the next handler in the chain
	}
}
