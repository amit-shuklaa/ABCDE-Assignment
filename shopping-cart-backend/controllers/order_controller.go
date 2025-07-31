package controllers

import (
	"log" // Added for logging warnings
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"shopping-cart-backend/database"
	"shopping-cart-backend/models"
)

// CreateOrder handles POST /orders to convert an active cart into an order[cite: 16, 20].
func CreateOrder(c *gin.Context) {
	userID := c.MustGet("userID").(uint) // Get userID from authenticated context

	// Find the user's active cart
	var cart models.Cart
	// Ensure we only convert 'active' carts, not already 'ordered' ones
	if err := database.DB.Where("user_id = ? AND status = ?", userID, "active").First(&cart).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusBadRequest, gin.H{"error": "No active cart found to convert to order"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error finding active cart"})
		return
	}

	// Check if the cart has items before creating an order
	var cartItems []models.CartItem
	database.DB.Where("cart_id = ?", cart.ID).Find(&cartItems)
	if len(cartItems) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot place an order with an empty cart"})
		return
	}

	// Create the order
	order := models.Order{
		CartID:    cart.ID,
		UserID:    userID,
		CreatedAt: time.Now(),
	}

	if err := database.DB.Create(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
		return
	}

	// Update the cart status to 'ordered' after conversion [cite: 16]
	if err := database.DB.Model(&cart).Update("status", "ordered").Error; err != nil {
		// Log this error, but don't necessarily return failure as order is already created
		log.Printf("Warning: Failed to update cart status for cart_id %d: %v", cart.ID, err)
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order placed successfully", "order_id": order.ID, "cart_id": order.CartID})
}

// GetOrdersByUserID handles GET /orders to list all orders for the authenticated user[cite: 17, 20].
func GetOrdersByUserID(c *gin.Context) {
	userID := c.MustGet("userID").(uint) // Get userID from authenticated context

	var orders []models.Order
	// Fetch all orders associated with the authenticated user
	if err := database.DB.Where("user_id = ?", userID).Find(&orders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error retrieving orders"})
		return
	}

	var orderIDs []uint
	for _, order := range orders {
		orderIDs = append(orderIDs, order.ID)
	}

	c.JSON(http.StatusOK, gin.H{"orders": orders, "order_ids": orderIDs}) // Return individual IDs as well for the UI requirement
}
