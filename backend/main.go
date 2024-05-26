package main

import (
	"go-gin-backend/database"
	"go-gin-backend/routes"
)

func main() {
	database.Connect()

	r := routes.SetupRouter()
	r.Run(":8080") // Run on port 8080
}
