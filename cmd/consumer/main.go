package main

import (
	"log"

	"calendar-service/config"
	"calendar-service/db"
	"calendar-service/models"
	"calendar-service/sqs"
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
	// Load configuration
	sqlConfig, err := config.LoadSQLConfig()
	if err != nil {
		log.Fatalf("Failed to load SQL config: %v", err)
	}

	if err := db.Init(sqlConfig); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	migrate()
	sqsConfig, err := config.LoadSQSConfig()
	if err != nil {
		log.Fatalf("Failed to load SQS config: %v", err)
	}

	// Initialize SQS client
	sqs.InitSQSClient()

	// Start SQS consumer
	sqs.StartPollingQueue(sqsConfig.Queues["calendar-entry"])
}
