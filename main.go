package main

import (
	"events-management/database"
	"events-management/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
}
