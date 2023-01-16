package routes

import (
	"rest-api/controllers"

	"github.com/gin-gonic/gin"
)

func addPostRoutes(rg *gin.RouterGroup) {
	posts := rg.Group("/posts")

	posts.GET("/", controllers.FindBooks)
	posts.POST("/create", controllers.CreatePost)
}
