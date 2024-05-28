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
	r.POST("/educations", controllers.CreateEducation)
	r.GET("/educations", controllers.GetEducations)
	r.GET("/educations/:id", controllers.GetEducation)
	r.PUT("/educations/:id", controllers.UpdateEducation)
	r.DELETE("/educations/:id", controllers.DeleteEducation)

	//Group routes
	r.POST("/groups", controllers.CreateGroup)
	r.GET("/groups", controllers.GetGroups)
	r.GET("/groups/:id", controllers.GetGroup)
	r.PUT("/groups/:id", controllers.UpdateGroup)
	r.DELETE("/groups/:id", controllers.DeleteGroup)

	return r
}
