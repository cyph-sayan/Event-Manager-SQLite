package routes

import (
	"events-management/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/getEvents", getEvents)
	server.GET("/getEvent/:id", getEvent)

	authenticated := server.Group("/")
	authenticated.Use(middleware.Authorization)
	authenticated.POST("/createEvent", saveEvent)
	authenticated.PUT("/event/:id", updateEvent)
	authenticated.DELETE("/deleteEvent/:id", deleteEvent)

	server.POST("/createUser", createUser)
	server.POST("/login", loginUser)
	server.Run(":8080")
}
