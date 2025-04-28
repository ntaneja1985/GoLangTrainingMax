package routes

import (
	"example.com/rest-api/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	//Register a handler with the Gin Engine
	//First argument is the path to which request is sent
	//This is similar to Minimal API we have .NET Core
	//Second argument can be an anonymous function, or it can be a named function
	server.GET("/events", getEvents) //GET,PUT,POST,PATCH,DELETE

	//Create a grouping of routes
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", middlewares.Authenticate, createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	authenticated.POST("/events/:id/register", registerForEvent)     //Register for an event
	authenticated.DELETE("/events/:id/register", cancelRegistration) //Cancelling the registration

	//Approach 1: Here the middlewares are executed from left to right
	//server.POST("/events", middlewares.Authenticate, createEvent)
	server.GET("/events/:id", getEvent)
	//server.PUT("/events/:id", updateEvent)
	//server.DELETE("/events/:id", deleteEvent)
	server.POST("/signup", signup)
	server.POST("/login", login)
}
