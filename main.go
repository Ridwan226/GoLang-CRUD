package main

import (
	"belajar/config"
	"belajar/controllers"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	
	err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }
	
	config.ConnectDB()
	
	server := gin.Default()

	// API Routes
	api := server.Group("/api")
	{
		api.POST("/events", controllers.CreateEvent)
		api.GET("/events", controllers.GetAllEvents)
		api.GET("/event/:id", controllers.GetEventById)
		api.PUT("/event/:id", controllers.UpdateEventById)
		api.DELETE("/event/:id", controllers.DeleteEventById)
		
		
		api.POST("/register", controllers.RegisterUser)
		
	}

	server.Run(":8080")
}


