package routes

import (
	"rest-api/controllers"

	"github.com/gin-gonic/gin"
)

func addUserRoutes(rg *gin.RouterGroup) {
	users := rg.Group("/users")

	users.POST("/register", controllers.SignUp())
	users.POST("/login", controllers.Login())
}
