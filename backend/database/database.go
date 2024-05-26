package database

import (
	"go-gin-backend/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	dsn := "blueberry:BL@2024er_@tcp(localhost:3306)/blueberry?charset=utf8mb4&parseTime=True&loc=Local"

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database! ")
	}

	err = database.AutoMigrate(&models.User{}, &models.Feed{})

	if err != nil {
		panic("Failed to migrate database!")
	}

	DB = database

}
