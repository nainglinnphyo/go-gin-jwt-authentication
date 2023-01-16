package routes

import (
	"rest-api/controllers"

	"github.com/gin-gonic/gin"
)

func addUserRoutes(rg *gin.RouterGroup) {
	users := rg.Group("/users")

	users.POST("/create", controllers.CreateUser())
	users.GET("/", controllers.GetAllUsers())
}
