package database

import (
	"fmt"
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

	fmt.Println("Database Connected Successfully")

	err = database.AutoMigrate(&models.User{}, &models.Feed{}, &models.Education{}, &models.Group{})

	if err != nil {
		panic("Failed to migrate database!")
	}

	fmt.Println("Database Migration Successfully")

	DB = database

}
