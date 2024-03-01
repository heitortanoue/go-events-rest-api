package main

import (
	"heitortanoue/rest-api/db"
	"heitortanoue/rest-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default() // Create a new server that uses the default middleware
	db.InitDB()

	routes.RegisterRoutes(server)
	server.Run(":8080") // Listen and serve on port 8080
}
