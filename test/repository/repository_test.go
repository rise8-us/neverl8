package repository_test

import (
	"os"
	"testing"
	"time"

	"github.com/rise8-us/neverl8/model"
	"github.com/rise8-us/neverl8/repository"
	testutil "github.com/rise8-us/neverl8/test/testcontainers"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var db *gorm.DB

// TestMain sets up the test database
func TestMain(m *testing.M) {
	testDB := testutil.SetupTestDB()
	db = testDB.DB

	code := runTests(m, testDB.TearDown)
	os.Exit(code)
}

// While the m.Run() function already runs all the tests, os.Exit() does not respect the defer method.
// This function is used to ensure that the test database is torn down after all tests are run.
func runTests(m *testing.M, tearDown func()) int {
	defer tearDown()
	return m.Run()
}

func GetSampleMeeting() (*model.Meetings, *[]model.Hosts) {
	// Create new Hosts
	hosts := &[]model.Hosts{
		{HostName: "Host 1"},
		{HostName: "Host 2"},
	}

	meetingDuration := 60
	// New Meeting to be created
	meeting := &model.Meetings{
		CandidateID: 2,
		Calendar:    "Example Calendar",
		Duration:    60,
		Title:       "Example Session",
		Description: "Discuss the future of NeverL8",
		HasBotGuest: false,
		StartTime:   time.Now(),
		EndTime:     time.Now().Add(time.Minute * time.Duration(meetingDuration)),
	}

	return meeting, hosts
}

func TestCreateMeeting(t *testing.T) {
	repo := repository.NewMeetingRepository(db)

	meeting, hosts := GetSampleMeeting()

	// Create Meeting
	createdMeeting, err := repo.CreateMeeting(meeting, *hosts)
	assert.NoError(t, err, "expected no error")
	assert.NotNil(t, createdMeeting, "expected meeting to be created")
	assert.Equal(t, uint(1), meeting.ID, "expected meeting id to be 1")
	assert.Equal(t, meeting.CandidateID, createdMeeting.CandidateID, "expected candidate id to match")
	assert.Equal(t, createdMeeting.Calendar, meeting.Calendar, "expected calendar to match")
	assert.Equal(t, createdMeeting.Duration, meeting.Duration, "expected duration to match")
	assert.Equal(t, createdMeeting.Title, meeting.Title, "expected title to match")
	assert.Equal(t, createdMeeting.Description, meeting.Description, "expected description to match")
	assert.Equal(t, createdMeeting.HasBotGuest, meeting.HasBotGuest, "expected has bot guest to match")
	assert.Equal(t, createdMeeting.StartTime, meeting.StartTime, "expected start time to match")
	assert.Equal(t, createdMeeting.EndTime, meeting.EndTime, "expected end time to match")
}

func TestGetAllMeetings(t *testing.T) {
	repo := repository.NewMeetingRepository(db)

	// Get all meetings
	meetings, err := repo.GetAllMeetings()
	assert.NoError(t, err, "expected no error")
	assert.NotNil(t, meetings, "expected meetings to be retrieved")
	assert.Equal(t, 1, len(meetings), "expected 1 meeting to be retrieved")
}
