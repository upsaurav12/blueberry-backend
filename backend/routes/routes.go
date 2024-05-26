package routes

import (
	"go-gin-backend/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// User routes
	r.POST("/users", controllers.CreateUser)
	r.GET("/users", controllers.GetUsers)
	r.GET("/users/:id", controllers.GetUser)
	r.PUT("/users/:id", controllers.UpdateUser)
	r.DELETE("/users/:id", controllers.DeleteUser)

	// Feed routes
	r.POST("/feeds", controllers.CreateFeed)
	r.GET("/feeds", controllers.GetFeeds)
	r.GET("/feeds/:id", controllers.GetFeed)
	r.PUT("/feeds/:id", controllers.UpdateFeed)
	r.DELETE("/feeds/:id", controllers.DeleteFeed)

	return r
}
