package calendar_test

import (
	"calendar-service/config"
	"calendar-service/db"
	"calendar-service/handlers"
	"calendar-service/models"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func setupConfig() error {
    viper.SetConfigName("config") // config.yaml without extension
    viper.AddConfigPath("..")     // Go one directory up from the tests directory
    viper.SetConfigType("yaml")

    if err := viper.ReadInConfig(); err != nil {
        return err
    }

    return nil
}

func setupRouter() *gin.Engine {
    r := gin.Default()
    r.GET("/calendar/active", handlers.GetActiveCalendarEntries)
    return r
}

func TestGetActiveCalendarEntriesStatus(t *testing.T) {
    // Set up configuration
    if err := setupConfig(); err != nil {
        t.Fatalf("Failed to set up config: %v", err)
    }

    // Initialize database
    sqlConfig, err := config.LoadSQLConfig()
    if err != nil {
        t.Fatalf("Failed to load SQL config: %v", err)
    }
    if err := db.Init(sqlConfig); err != nil {
        t.Fatalf("Failed to initialize database: %v", err)
    }

    // Migrate database schema
    db.DB.AutoMigrate(&models.CalendarEntry{}, &models.LastProcessedTime{})

    // Insert a calendar entry
    entry := models.CalendarEntry{
        StartDate: time.Now().Add(-1 * time.Hour),
        StopDate:  time.Now().Add(1 * time.Hour),
        CreatedAt: time.Now(),
        UpdatedAt: time.Now(),
    }
    db.DB.Create(&entry)

    // Set up router
    r := setupRouter()

    // Perform the request
    req, _ := http.NewRequest("GET", "/calendar/active", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    // Assert that the response status is OK
    assert.Equal(t, http.StatusOK, w.Code)
}
