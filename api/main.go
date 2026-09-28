package main

import (
	"net/http"

	"com.example/event/models"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.GET("/events", getEvents)
	server.POST("/events", createEvent)
	server.Run(":8080") //localhost:8080
}

func getEvents(context *gin.Context) {
	events := models.GetAllEvents()
	//context.JSON(http.StatusOK, gin.H{"message": "Hello!"})
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil { //binds request json to struct
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	event.ID = 1
	event.UserID = 1
	event.Save()
	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}
