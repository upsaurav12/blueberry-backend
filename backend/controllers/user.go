package controllers

import (
	"go-gin-backend/database"
	"go-gin-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateUser handles creating a new user
func CreateUser(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		UserImage:    input.UserImage,
		FirstName:    input.FirstName,
		LastName:     input.LastName,
		Summary:      input.Summary,
		Email:        input.Email,
		Location:     input.Location,
		Role:         input.Role,
		PhoneNumber:  input.PhoneNumber,
		ActiveStatus: input.ActiveStatus,
		LinkedIn:     input.LinkedIn,
		Instagram:    input.Instagram,
		Password:     input.Password,
	}
	database.DB.Create(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// GetUser handles fetching a single user by ID
func GetUser(c *gin.Context) {
	var user models.User
	if err := database.DB.Preload("Feeds").Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// GetUsers handles fetching all users
func GetUsers(c *gin.Context) {
	var users []models.User
	database.DB.Preload("Feeds").Find(&users)

	c.JSON(http.StatusOK, gin.H{"data": users})
}

// UpdateUser handles updating a user
func UpdateUser(c *gin.Context) {
	var user models.User
	if err := database.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found!"})
		return
	}

	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Model(&user).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": user})
}

// DeleteUser handles deleting a user
func DeleteUser(c *gin.Context) {
	var user models.User
	if err := database.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found!"})
		return
	}

	database.DB.Delete(&user)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
