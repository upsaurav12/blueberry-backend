package controllers

import (
	"net/http"
	"time"

	"go-gin-backend/database"
	"go-gin-backend/models"

	"github.com/gin-gonic/gin"
)

// CreateFeedInput defines the structure for feed creation input
type CreateFeedInput struct {
	UserID      uint   `json:"user_id" binding:"required"`
	TextContent string `json:"text_content" binding:"required"`
	Image       string `json:"image"`
	ProfileURL  string `json:"profile_url"`
}

// CreateFeed handles creating a new feed
func CreateFeed(c *gin.Context) {
	var input CreateFeedInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate user existence
	var user models.User
	if err := database.DB.Where("id = ?", input.UserID).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found!"})
		return
	}

	feed := models.Feed{
		UserID:      input.UserID,
		TextContent: input.TextContent,
		Timestamp:   time.Now(),
		Image:       input.Image,
		ProfileURL:  input.ProfileURL,
	}

	database.DB.Create(&feed)
	c.JSON(http.StatusOK, gin.H{"data": feed})
}

// GetFeeds handles fetching all feeds
func GetFeeds(c *gin.Context) {
	var feeds []models.Feed
	database.DB.Find(&feeds)

	c.JSON(http.StatusOK, gin.H{"data": feeds})
}

// GetFeed handles fetching a single feed by ID
func GetFeed(c *gin.Context) {
	var feed models.Feed
	if err := database.DB.Where("id = ?", c.Param("id")).First(&feed).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Feed not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": feed})
}

// UpdateFeed handles updating a feed
func UpdateFeed(c *gin.Context) {
	var feed models.Feed
	if err := database.DB.Where("id = ?", c.Param("id")).First(&feed).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Feed not found!"})
		return
	}

	var input models.Feed
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Model(&feed).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": feed})
}

// DeleteFeed handles deleting a feed
func DeleteFeed(c *gin.Context) {
	var feed models.Feed
	if err := database.DB.Where("id = ?", c.Param("id")).First(&feed).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Feed not found!"})
		return
	}

	database.DB.Delete(&feed)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
