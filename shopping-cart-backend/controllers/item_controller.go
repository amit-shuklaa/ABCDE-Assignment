package controllers

import (
	"net/http"
	"time"

	"shopping-cart-backend/database"
	"shopping-cart-backend/models"

	"github.com/gin-gonic/gin"
)

// Input struct for item creation
type CreateItemInput struct {
	Name   string `json:"name" binding:"required"`
	Status string `json:"status"` // Optional, default to "available"
}

// CreateItem handles POST /items to create a new item[cite: 20].
func CreateItem(c *gin.Context) {
	var input CreateItemInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item := models.Item{
		Name:      input.Name,
		Status:    "available", // Default status
		CreatedAt: time.Now(),
	}
	if input.Status != "" {
		item.Status = input.Status
	}

	if result := database.DB.Create(&item); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item created successfully", "item": item})
}

// GetItems handles GET /items to list all available items[cite: 17, 20].
func GetItems(c *gin.Context) {
	var items []models.Item
	if result := database.DB.Find(&items); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"items": items})
}
