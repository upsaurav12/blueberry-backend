/*package database

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
*/

package database

import (
	"fmt"
	"go-gin-backend/models"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	// Load the environment variables (assuming you have set them)
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// Construct the DSN
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)

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
