package controllers

import (
	"go-gin-backend/database"
	"go-gin-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateGroupInput struct {
	GroupId       uint   `json:"group_id" gorm:"primary_key"`
	UserID        uint   `json:"user_id" binding:"required"`
	GroupName     string `json:"group_name"`
	AdminUserId   uint   `json:"admin_id"`
	IsActive      bool   `json:"isactive"`
	Description   string `json:"description"`
	GroupImageURL string `json:"groupimage"`
}

func CreateGroup(c *gin.Context) {
	var input CreateGroupInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Check the existence of the user
	var user models.User
	if err := database.DB.Where("id = ?", input.UserID).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found!"})
		return
	}

	group := models.Group{
		UserID:        input.UserID,
		GroupId:       input.GroupId,
		GroupName:     input.GroupName,
		AdminUserId:   input.AdminUserId,
		IsActive:      input.IsActive,
		Description:   input.Description,
		GroupImageURL: input.GroupImageURL,
	}

	database.DB.Create(&group)
	c.JSON(http.StatusOK, gin.H{"data": group})

}

func GetGroups(c *gin.Context) {
	var groups []models.Group
	database.DB.Find(&groups)

	c.JSON(http.StatusOK, gin.H{"data": groups})
}

func GetGroup(c *gin.Context) {

}

func UpdateGroup(c *gin.Context) {

}

func DeleteGroup(c *gin.Context) {

}
