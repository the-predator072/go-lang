package routes

import (
	"net/http"
	"strconv"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func getEvents(ctx *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't fetch events."})
		return
	}
	ctx.JSON(http.StatusOK, events)
}

func createEvent(ctx *gin.Context) {
	var event models.Event
	err := ctx.ShouldBindJSON(&event)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Couldn't parse the request body."})
		return
	}
	event.UserId = 1
	err = event.Save()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't create the event."})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "Event Created!", "event": event})
}

func getEvent(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Couldn't parse event id."})
		return
	}
	event, err := models.GetEventById(eventId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't fetch the event."})
		return
	}
	ctx.JSON(http.StatusOK, event)
}

func updateEvent(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Couldn't parse event ID."})
		return
	}

	_, err = models.GetEventById(eventId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't fetch the event."})
		return
	}

	var updatedEvent models.Event

	err = ctx.ShouldBindJSON(&updatedEvent)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Couldn't parse the request body."})
		return
	}

	updatedEvent.ID = eventId
	err = updatedEvent.Update()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't update event."})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Updated Successfully.", "data": updatedEvent})
}

func deleteEvent(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Couldn't parse event ID."})
		return
	}

	event, err := models.GetEventById(eventId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't fetch the event."})
		return
	}
	err = event.Delete()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't delete event."})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully"})
}
