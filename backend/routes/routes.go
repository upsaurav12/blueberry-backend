package routes

import (
	"go-gin-backend/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// User routes
	r.GET("/users", controllers.GetUsers)
	r.GET("/users/:id", controllers.GetUser)
	r.POST("/users", controllers.CreateUser)
	r.PUT("/users/:id", controllers.UpdateUser)
	r.DELETE("/users/:id", controllers.DeleteUser)

	// Feed routes
	r.POST("/feeds", controllers.CreateFeed)
	r.GET("/feeds", controllers.GetFeeds)
	r.GET("/feeds/:id", controllers.GetFeed)
	r.PUT("/feeds/:id", controllers.UpdateFeed)
	r.DELETE("/feeds/:id", controllers.DeleteFeed)

	//Education routes
	r.POST("/education", controllers.CreateEducation)
	r.GET("/educations", controllers.GetEducations)
	r.GET("/education/:id", controllers.GetEducation)
	r.PUT("/education/:id", controllers.UpdateEducation)
	r.DELETE("/education/:id", controllers.DeleteEducation)

	//Group routes
	r.POST("/group", controllers.CreateGroup)
	r.GET("/groups", controllers.GetGroups)
	r.GET("/group/:id", controllers.GetGroup)
	r.PUT("/group/:id", controllers.UpdateGroup)
	r.DELETE("/group/:id", controllers.DeleteGroup)

	return r
}
