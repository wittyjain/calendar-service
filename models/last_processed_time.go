// models/last_processed_time.go
package models

import (
	"time"
)

type LastProcessedTime struct {
	ID             int       `gorm:"primaryKey"`
	LastProcessed  time.Time `gorm:"type:datetime(3)"`
}
