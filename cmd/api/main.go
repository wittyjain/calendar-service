package main

import (
	"log"

	"calendar-service/config"
	"calendar-service/db"
	"calendar-service/handlers"
	"calendar-service/middleware"
	"calendar-service/models"
	"calendar-service/sqs"

	"github.com/gin-gonic/gin"
)

func migrate() {
    err := db.DB.AutoMigrate(
        &models.CalendarEntry{},
		&models.LastProcessedTime{},
    )
    if err != nil {
        log.Fatalf("migration failed: %v", err)
    }
}


func main() {
	// Load SQL configuration
	sqlConfig, err := config.LoadSQLConfig()
	if err != nil {
		log.Fatalf("Failed to load SQL config: %v", err)
	}

	// Initialize the database
	if err := db.Init(sqlConfig); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	//TODO: Remove this in production
	migrate()

	sqs.InitSQSClient()
	//Gin router
	router := gin.Default()
	router.Use(middleware.LoggingMiddleware())
	router.Use(middleware.ExceptionMiddleware())
	
	// API routes	
	router.POST("/calendar", handlers.CreateCalendarEntry)
	router.GET("/calendar/active", handlers.GetActiveCalendarEntries)
	router.POST("/process-calendar-entries", handlers.ProcessCalendarEntriesHandler)
	router.Run(":8080")
}
