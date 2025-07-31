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

// Input struct for adding items to cart
type AddToCartInput struct {
	ItemID uint `json:"item_id" binding:"required"`
}

// AddItemsToCart handles POST /carts to create a cart or add items to an existing one.
func AddItemsToCart(c *gin.Context) {
	userID := c.MustGet("userID").(uint) // Get userID from authenticated context

	var input AddToCartInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if item exists
	var item models.Item
	if err := database.DB.First(&item, input.ItemID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error checking item"})
		return
	}

	// A single user can have only a single cart
	var cart models.Cart
	result := database.DB.Where("user_id = ?", userID).First(&cart)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			// Create a new cart if not found
			cart = models.Cart{
				UserID:    userID,
				Name:      "Shopping Cart",
				Status:    "active", // Mark as active
				CreatedAt: time.Now(),
			}
			if err := database.DB.Create(&cart).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create cart"})
				return
			}
			// Update User's CartID (optional, but good for relationship tracking)
			database.DB.Model(&models.User{}).Where("id = ?", userID).Update("cart_id", cart.ID)
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error finding cart"})
			return
		}
	}

	// Add item to cart_items (many-to-many relationship)
	cartItem := models.CartItem{
		CartID: cart.ID,
		ItemID: input.ItemID,
	}

	// Check if item already exists in cart for simplicity (prevents duplicate entries if not managed by quantity)
	var existingCartItem models.CartItem
	res := database.DB.Where("cart_id = ? AND item_id = ?", cart.ID, input.ItemID).First(&existingCartItem)
	if res.Error == nil {
		c.JSON(http.StatusOK, gin.H{"message": "Item already in cart"})
		return
	} else if res.Error != gorm.ErrRecordNotFound {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error checking cart item"})
		return
	}

	if err := database.DB.Create(&cartItem).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add item to cart"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item added to cart successfully", "cart_id": cart.ID, "item_id": input.ItemID})
}

// GetCartByUserID handles GET /carts to list the authenticated user's cart and its items.
// It assumes a user has one "active" cart or the latest created one.
func GetCartByUserID(c *gin.Context) {
	userID := c.MustGet("userID").(uint) // Get userID from authenticated context

	var cart models.Cart
	// Eager load CartItems (the join table entries)
	// We're filtering for an 'active' cart, or you might choose the most recent.
	if err := database.DB.Where("user_id = ? AND status = ?", userID, "active").Preload("Items").First(&cart).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"message": "No active cart found for this user"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error retrieving cart"})
		return
	}

	// To display item names, you would typically fetch item details for each CartItem.
	// For simplicity, as per requirement, we'll return cart_id and item_id.
	var cartItemsDetails []gin.H
	for _, ci := range cart.Items {
		// Optionally, fetch item name for better display
		var item models.Item
		if err := database.DB.First(&item, ci.ItemID).Error; err != nil {
			log.Printf("Warning: Could not find item %d for cart %d: %v", ci.ItemID, ci.CartID, err)
			cartItemsDetails = append(cartItemsDetails, gin.H{"cart_id": ci.CartID, "item_id": ci.ItemID, "item_name": "Unknown Item"})
		} else {
			cartItemsDetails = append(cartItemsDetails, gin.H{"cart_id": ci.CartID, "item_id": ci.ItemID, "item_name": item.Name})
		}
	}

	c.JSON(http.StatusOK, gin.H{"cart": gin.H{
		"id":         cart.ID,
		"user_id":    cart.UserID,
		"status":     cart.Status,
		"items":      cartItemsDetails,
		"created_at": cart.CreatedAt,
	}})
}
