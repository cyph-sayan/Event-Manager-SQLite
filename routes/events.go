package routes

import (
	"events-management/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, map[string]string{"message": "error"})
		return
	}
	context.JSON(http.StatusOK, events)
	return
}

func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, map[string]string{"message": "error"})
		return
	}

	event, err := models.GetEventById(int(eventId))
	if err != nil {
		fmt.Print(err)
		context.JSON(http.StatusBadRequest, map[string]string{"message": "error"})
		return
	}

	context.JSON(http.StatusOK, event)
	return
}

func saveEvent(context *gin.Context) {
	userId := context.GetInt("userId")
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request"})
		return
	}
	event.UserId = int(userId)
	event.Save()
	context.JSON(http.StatusCreated, gin.H{"message": "Event created successfully", "event": event})
	return
}

func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, map[string]string{"message": "error"})
		return
	}
	retrievedData, err := models.GetEventById(int(eventId))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the event"})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse JSON"})
		return
	}

	userId := int(context.GetInt64("userId"))
	if userId != retrievedData.UserId {
		context.JSON(http.StatusForbidden, gin.H{"message": "User not allowed to update event"})
		return
	}
	updatedEvent.Id = eventId
	err = updatedEvent.UpdateEvent()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update event"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event updated successfully"})
}

func deleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, map[string]string{"message": "error"})
		return
	}
	var event models.Event
	event, err = models.GetEventById(int(eventId))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the event"})
		return
	}

	err = event.DeleteEvent()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete the event"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event updated successfully"})
}
