package handlers

import (
	"calendar-service/sqs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ProcessCalendarEntriesHandler(c *gin.Context) {
	err := sqs.ProcessCalendarEntryProducer()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
