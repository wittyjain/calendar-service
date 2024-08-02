package repositories

import (
	"calendar-service/db"
	"calendar-service/models"
	"log"
	"time"

	"gorm.io/gorm"
)

const defaultTime = "1970-01-01T00:00:00Z"


func FetchLastProcessedTime() (time.Time, error) {
    var lastProcessed models.LastProcessedTime
    result := db.DB.First(&lastProcessed)
    if result.Error != nil {
        if result.Error == gorm.ErrRecordNotFound {
            log.Println("Record not found, returning default time")
            defaultParsedTime, err := time.Parse(time.RFC3339, defaultTime)
            if err != nil {
                log.Printf("Failed to parse default time: %v", err)
                return time.Time{}, err
            }
            return defaultParsedTime, nil
        }
        log.Printf("Error fetching last processed time: %v", result.Error)
        return time.Time{}, result.Error
    }
    return lastProcessed.LastProcessed, nil
}
func UpdateLastProcessedTime(newTime time.Time) error {
    var lastProcessedTime models.LastProcessedTime

    // Check if there's already an entry in the table
    err := db.DB.First(&lastProcessedTime).Error
    if err != nil {
        if err == gorm.ErrRecordNotFound {
            // If no record exists, create a new one
            lastProcessedTime.LastProcessed = newTime
            log.Printf("Creating new record with time: %v", newTime)
            return db.DB.Create(&lastProcessedTime).Error
        }
        return err
    }

    // Update the existing record with the new time
    lastProcessedTime.LastProcessed = newTime
    log.Printf("Updating record with new time: %v", newTime)
    return db.DB.Save(&lastProcessedTime).Error
}
