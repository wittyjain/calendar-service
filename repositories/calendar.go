package repositories

import (
	"calendar-service/db"
	"calendar-service/models"
	"time"
)



func FetchCalendarEntriesSince(since time.Time) ([]models.CalendarEntry, error) {
	var entries []models.CalendarEntry
	now := time.Now()
	err := db.DB.Where("stop_date > ? AND updated_at > ?", now, since).Find(&entries).Error
	if err != nil {
		return nil, err
	}
	return entries, nil
}

func CreateCalendarEntry(entry *models.CalendarEntry) error {
	entry.CreatedAt = time.Now()
	entry.UpdatedAt = time.Now()
	return db.DB.Create(entry).Error
}

// GetActiveCalendarEntries retrieves all active calendar entries.
func GetActiveCalendarEntries(now time.Time) ([]models.CalendarEntry, error) {
	var entries []models.CalendarEntry
	err := db.DB.Where("stop_date >= ?", now).Find(&entries).Error
	return entries, err
}