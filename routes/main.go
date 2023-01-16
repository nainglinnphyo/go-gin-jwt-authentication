package routes

import (
	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func Run() {
	getRoutes()
	router.Run(":5050")
}

func getRoutes() {
	apiRoutes := router.Group("/api")
	addUserRoutes(apiRoutes)
	addPostRoutes(apiRoutes)
}
