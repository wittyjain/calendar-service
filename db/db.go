package db

import (
	"calendar-service/config"
	"calendar-service/models"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init(cfg *config.SQLConfig) error {
	dsn := cfg.Username + ":" + cfg.Password + "@tcp(" + cfg.URL + ")/" + cfg.Database + "?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	for retries := 0; retries < 10; retries++ {
		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}

		log.Println("Database not ready, retrying...")
		time.Sleep(5 * time.Second)
	}
	
	if err != nil {
		return err
	}

	// Auto migrate the CalendarEntry model
	if err := DB.AutoMigrate(&models.CalendarEntry{}); err != nil {
		return err
	}

	log.Println("Database connection initialized and tables migrated")
	return nil
}
