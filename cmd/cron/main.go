// main.go
package main

import (
	"calendar-service/config"
	"calendar-service/db"
	"calendar-service/models"
	"calendar-service/sqs"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/robfig/cron/v3"
)

func migrate() {
	err := db.DB.AutoMigrate(
		&models.CalendarEntry{}, // Add more models as needed
	)
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
}

func main() {
	sqlConfig, err := config.LoadSQLConfig()
	if err != nil {
		log.Fatalf("Failed to load SQL config: %v", err)
	}

	if err := db.Init(sqlConfig); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	migrate()

	// Initialize SQS client
	sqs.InitSQSClient()

	// Create a new cron scheduler
	c := cron.New()

	// Schedule the job
	c.AddFunc("@every 1m", func() {
		// Call the producer function
		err := sqs.ProcessCalendarEntryProducer()
		if err != nil {
			log.Printf("Error processing calendar entries: %v", err)
		}
	})

	// Start the cron scheduler
	c.Start()
	log.Println("Cron scheduler started.")

	// Set up channel to listen for interrupt signals
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	// Block until a signal is received
	<-stop
	log.Println("Received shutdown signal, exiting...")
	c.Stop() // Stop the cron scheduler gracefully
}
