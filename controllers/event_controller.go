package controllers

import (
	"belajar/config"
	"belajar/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event data", "err": err.Error()})
		return
	}

	event.UserID = 1
	
	config.DB.Create(&event)

	context.JSON(http.StatusCreated, gin.H{"message": "Event created successfully", "event": event})
	
}


func GetAllEvents(c *gin.Context) {
	var events []models.Event
	config.DB.Find(&events)
	c.JSON(http.StatusOK, gin.H{"status": "success", "events": events})
}

func GetEventById(c *gin.Context) {
	var event models.Event
	paramsId := c.Param("id")
	if paramsId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Event ID is required"})
		return
	}
	var evenData = config.DB.First(&event, paramsId).Error;
	if evenData != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Event not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "event": event})
}

func UpdateEventById(c *gin.Context) {
	var event models.Event
	paramsId := c.Param("id")
	if paramsId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Event ID is required"})
		return
	}
	var input models.Event
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event data", "err": err.Error()})
		return
	}
	
	config.DB.Model(&event).Updates(&input)
	c.JSON(http.StatusCreated, gin.H{"status": "Update success", "event": event})
}

func DeleteEventById(c *gin.Context){
	var event models.Event
	
	paramsId := c.Param("id")
	
	var evenData = config.DB.First(&event, paramsId).Error;
	if evenData != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Event not found"})
		return
	}

	config.DB.Unscoped().Delete(&event)
	c.JSON(http.StatusOK, gin.H{"status": "Delete success", "event": event})
}