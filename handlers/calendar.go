// handlers/calendar_handler.go
package handlers

import (
	"calendar-service/models"
	"calendar-service/repositories"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// CreateCalendarEntry handles the creation of a new calendar entry.
func CreateCalendarEntry(c *gin.Context) {
	var entry models.CalendarEntry
	if err := c.BindJSON(&entry); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := repositories.CreateCalendarEntry(&entry); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, entry)
}

// GetActiveCalendarEntries retrieves and returns all active calendar entries.
func GetActiveCalendarEntries(c *gin.Context) {
	now := time.Now()
	entries, err := repositories.GetActiveCalendarEntries(now)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, entries)
}
