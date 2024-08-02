package models

import (
	"time"
)

type CalendarEntry struct {
	ID        uint      `gorm:"primaryKey"`
	StartDate time.Time `gorm:"not null"`
	StopDate  time.Time `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
