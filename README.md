# event-manager-crud-api
A simple and lightweight crud api for creating and managing events implemented using golang and gin. SQLite is used as the DB and endpoints are protected using JWT token.

The various endpoints are : 
1. GET /getEvents : Gets a list of events present in the DB.
2. GET /getEvent/:id : Gets a particular event according to the event id passed in request parameter

	authenticated := server.Group("/")
	authenticated.Use(middleware.Authorization)
	authenticated.POST("/createEvent", saveEvent)
	authenticated.PUT("/event/:id", updateEvent)
	authenticated.DELETE("/deleteEvent/:id", deleteEvent)

	server.POST("/createUser", createUser)
	server.POST("/login", loginUser)
