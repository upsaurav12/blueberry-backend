package controllers

import (
	"net/http"
	"time"

	"go-gin-backend/database"
	"go-gin-backend/models"

	"github.com/gin-gonic/gin"
)

type CreateEducationInput struct {
	UserID       uint       `json:"user_id" binding:"required"`
	Institution  string     `json:"institution" binding:"required"`
	Degree       string     `json:"degree" binding:"required"`
	FieldOfStudy string     `json:"field_of_study" binding:"required"`
	StartDate    time.Time  `json:"start_date" binding:"required"`
	EndDate      *time.Time `json:"end_date"`
}

func CreateEducation(c *gin.Context) {
	var input CreateEducationInput
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

	education := models.Education{
		UserID:       input.UserID,
		Institution:  input.Institution,
		Degree:       input.Degree,
		FieldOfStudy: input.FieldOfStudy,
		StartDate:    input.StartDate,
		EndDate:      *input.EndDate,
	}

	database.DB.Create(&education)
	c.JSON(http.StatusOK, gin.H{"data": education})
}

func GetEducations(c *gin.Context) {
	var educations []models.Education
	if err := database.DB.Find(&educations).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": educations})
}

func GetEducation(c *gin.Context) {
	var education models.Education
	if err := database.DB.Where("education_id = ?", c.Param("id")).First(&education).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": education})
}

func UpdateEducation(c *gin.Context) {
	var education models.Education
	if err := database.DB.Where("education_id = ?", c.Param("id")).First(&education).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var input CreateEducationInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedEducation := models.Education{
		UserID:       input.UserID,
		Institution:  input.Institution,
		Degree:       input.Degree,
		FieldOfStudy: input.FieldOfStudy,
		StartDate:    input.StartDate,
		EndDate:      *input.EndDate,
	}

	if err := database.DB.Model(&education).Updates(updatedEducation).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": education})
}

func DeleteEducation(c *gin.Context) {
	var education models.Education
	if err := database.DB.Where("education_id = ?", c.Param("id")).First(&education).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Delete(&education).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": true})
}
