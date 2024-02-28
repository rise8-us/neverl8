package repository_test

import (
	"fmt"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/rise8-us/neverl8/model"
	"github.com/rise8-us/neverl8/repository"
	"github.com/rise8-us/neverl8/service"
)

func TestCreateMeeting(t *testing.T) {
	// Create a new meeting repository with the mock DB connection
	mockDB := NewMockDB()
	repo := repository.NewMeetingRepository(mockDB)

	// Create a new meeting service with the repository
	service := service.NewMeetingService(repo)

	// Mock test data
	meeting := &model.Meeting{
		Calendar:    "Test Calendar",
		Duration:    60,
		Title:       "Test Meeting",
		Description: "Test Description",
		Hosts:       "Test Hosts",
		HasBotGuest: false,
	}

	// Call the CreateMeeting function
	createdMeeting, err := service.CreateMeeting(meeting)

	// Assert that no error occurred
	assert.NoError(t, err, "expected no error")

	// Assert that the created meeting matches the input meeting
	assert.NotNil(t, createdMeeting, "expected meeting to be created")
	assert.Equal(t, meeting.Calendar, createdMeeting.Calendar, "expected calendar to match")
	assert.Equal(t, meeting.Duration, createdMeeting.Duration, "expected duration to match")
	assert.Equal(t, meeting.Title, createdMeeting.Title, "expected title to match")
	assert.Equal(t, meeting.Description, createdMeeting.Description, "expected description to match")
	assert.Equal(t, meeting.Hosts, createdMeeting.Hosts, "expected hosts to match")
	assert.Equal(t, meeting.HasBotGuest, createdMeeting.HasBotGuest, "expected hasBotGuest to match")
}

// NewMockDB creates a new instance the test db
func NewMockDB() *gorm.DB {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=postgres_test password=password sslmode=disable")
	if err != nil {
		panic(fmt.Sprintf("failed to connect database: %v", err))
	}
	db.AutoMigrate(&model.Meeting{})
	return db
}
