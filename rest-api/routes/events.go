// Include all request handler functions
package routes

import (
	"example.com/rest-api/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// This function will have an argument/pointer to the Gin Context similar to HttpContext
func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {

	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Fetch the userId from the context, more specifically the authentication middleware
	userId := context.GetInt64("userId")

	//event.ID = 1
	event.UserID = userId
	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err})
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event created", "event": event})
}

func getEvent(context *gin.Context) {
	fmt.Println(context.Param("id"))
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": event})
}

func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	//Fetch the userId from the context, more specifically the authentication middleware
	userId := context.GetInt64("userId")

	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	//Ensure that only the user who created the event should be able to update the event
	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "User does not have permission"})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err})
	}
	updatedEvent.ID = event.ID
	updatedEvent.UserID = event.UserID
	err = updatedEvent.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err})
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event updated", "event": updatedEvent})
}

func deleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err})
	}

	//Fetch the userId from the context, more specifically the authentication middleware
	userId := context.GetInt64("userId")

	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	//Ensure that only the user who created the event should be able to delete the event
	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "User does not have permission"})
		return
	}

	err = event.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err})
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event deleted", "event": event})
}
