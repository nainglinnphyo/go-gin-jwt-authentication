package main

import (
	"rest-api/config"
	"rest-api/routes"
)

func main() {
	// Our server will live in the routes package
	config.ConnectDB()
	routes.Run()
}
