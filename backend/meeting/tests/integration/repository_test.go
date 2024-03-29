package integration

import (
	"os"
	"testing"
	"time"

	"github.com/rise8-us/neverl8/meeting"
	"github.com/rise8-us/neverl8/model"
	testutil "github.com/rise8-us/neverl8/test_config"
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

func GetSampleHosts() []model.Host {
	hosts := &[]model.Host{
		{HostName: "Host 1", ID: 1, LastMeetingTime: time.Date(2021, time.January, 1, 0, 0, 0, 0, time.UTC)},
		{HostName: "Host 2", ID: 2, LastMeetingTime: time.Date(2021, time.January, 1, 0, 0, 0, 0, time.UTC)},
	}
	return *hosts
}

func GetSampleMeeting() *model.Meetings {
	// Create new Hosts
	hosts := GetSampleHosts()

	currentTime := time.Now().UTC().Truncate(time.Second)
	meetingDuration := 60
	// New Meeting to be created
	sampleMeeting := &model.Meetings{
		CandidateID: 2,
		Calendar:    "Example Calendar",
		Duration:    60,
		Title:       "Example Session",
		Description: "Discuss the future of NeverL8",
		HasBotGuest: false,
		StartTime:   currentTime,
		EndTime:     currentTime.Add(time.Minute * time.Duration(meetingDuration)),
		CreatedAt:   currentTime,
		Hosts:       hosts,
	}

	return sampleMeeting
}

func TestCreateMeeting(t *testing.T) {
	t.Cleanup(func() {
		db.Where("1 = 1").Delete(&model.Meetings{})
		db.Where("1 = 1").Delete(&model.Host{})
	})

	repo := meeting.NewMeetingRepository(db)

	sampleMeeting := GetSampleMeeting()

	// Create Meeting
	createdMeeting, err := repo.CreateMeeting(sampleMeeting)
	assert.NoError(t, err, "expected no error")
	assert.NotNil(t, createdMeeting, "expected meeting to be created")
	assert.Equal(t, uint(1), sampleMeeting.ID, "expected meeting id to be 1")
	assert.Equal(t, sampleMeeting.CandidateID, createdMeeting.CandidateID, "expected candidate id to match")
	assert.Equal(t, createdMeeting.Calendar, sampleMeeting.Calendar, "expected calendar to match")
	assert.Equal(t, createdMeeting.Duration, sampleMeeting.Duration, "expected duration to match")
	assert.Equal(t, createdMeeting.Title, sampleMeeting.Title, "expected title to match")
	assert.Equal(t, createdMeeting.Description, sampleMeeting.Description, "expected description to match")
	assert.Equal(t, createdMeeting.HasBotGuest, sampleMeeting.HasBotGuest, "expected has bot guest to match")
	assert.Equal(t, createdMeeting.StartTime, sampleMeeting.StartTime, "expected start time to match")
	assert.Equal(t, createdMeeting.EndTime, sampleMeeting.EndTime, "expected end time to match")
}

func TestGetAllMeetings(t *testing.T) {
	t.Cleanup(func() {
		db.Where("1 = 1").Delete(&model.Meetings{})
		db.Where("1 = 1").Delete(&model.Host{})
	})

	repo := meeting.NewMeetingRepository(db)

	sampleMeeting := GetSampleMeeting()

	// Create Meeting
	_, err := repo.CreateMeeting(sampleMeeting)
	assert.NoError(t, err, "expected no error")

	// Get all meetings
	meetings, err := repo.GetAllMeetings()
	assert.NoError(t, err, "expected no error")
	assert.NotNil(t, meetings, "expected meetings to be retrieved")
	assert.Equal(t, 1, len(meetings), "expected 1 meeting to be retrieved")
}

func TestGetMeetingByID(t *testing.T) {
	t.Cleanup(func() {
		db.Where("1 = 1").Delete(&model.Meetings{})
		db.Where("1 = 1").Delete(&model.Host{})
	})

	repo := meeting.NewMeetingRepository(db)

	sampleMeeting := GetSampleMeeting()

	// Create Meeting
	createdMeeting, err := repo.CreateMeeting(sampleMeeting)
	assert.NoError(t, err, "expected no error")

	// Get Meeting 1
	retrievedMeeting, err := repo.GetMeetingByID(createdMeeting.ID)
	assert.NoError(t, err, "expected no error")
	assert.NotNil(t, retrievedMeeting, "expected meeting to be retrieved")
	assert.Equal(t, createdMeeting.ID, retrievedMeeting.ID, "expected meetings to have the same IDs")
}

func TestCreateSampleMeeting(t *testing.T) {
	t.Cleanup(func() {
		db.Where("1 = 1").Delete(&model.SampleMeetings{})
	})

	repo := meeting.NewMeetingRepository(db)

	sampleMeeting := &model.SampleMeetings{Title: "Example Sample Meeting", Description: "Description of the Sample Meeting", Duration: 60}
	createdSampleMeeting := repo.CreateSampleMeeting(sampleMeeting)
	assert.NotNil(t, createdSampleMeeting, "expected sample meeting to be created")
	assert.Equal(t, uint(1), createdSampleMeeting.ID, "expected sample meeting id to be 1")
	assert.Equal(t, createdSampleMeeting, sampleMeeting, "expected sample meeting to equal retrieved sample meeting")
}

func TestGetSampleMeetings(t *testing.T) {
	t.Cleanup(func() {
		db.Where("1 = 1").Delete(&model.SampleMeetings{})
	})

	repo := meeting.NewMeetingRepository(db)

	sampleMeeting := &model.SampleMeetings{Title: "Example Sample Meeting", Description: "Description of the Sample Meeting", Duration: 60}
	createdSampleMeeting := repo.CreateSampleMeeting(sampleMeeting)
	assert.NotNil(t, createdSampleMeeting, "expected sample meeting to be created")

	sampleMeetings, err := repo.GetSampleMeetings()
	assert.NoError(t, err, "expected no error")
	assert.NotNil(t, sampleMeetings, "expected sample meetings to be retrieved")
	assert.Equal(t, 1, len(sampleMeetings), "expected 1 sample meeting to be retrieved")
}

// TestCreateMeetingWithSampleMeeting ensures meetings are properly created utilizing a sample meeting as a template.
func TestCreateMeetingWithSampleMeeting(t *testing.T) {
	t.Cleanup(func() {
		db.Where("1 = 1").Delete(&model.Meetings{})
		db.Where("1 = 1").Delete(&model.Host{})
	})

	repo := meeting.NewMeetingRepository(db)

	// Create a sample meeting
	sampleMeeting := &model.SampleMeetings{Title: "Example Sample Meeting", Description: "Description of the Sample Meeting", Duration: 60}
	createdSampleMeeting := repo.CreateSampleMeeting(sampleMeeting)
	assert.NotNil(t, createdSampleMeeting, "expected sample meeting to be created")

	// Create new Hosts
	hosts := GetSampleHosts()

	currentTime := time.Now().UTC().Truncate(time.Second)
	meetingDuration := 60
	// New Meeting to be created
	newMeeting := &model.Meetings{
		CandidateID: 2,
		Calendar:    "Example Calendar",
		Duration:    sampleMeeting.Duration, // Utilize sample meeting values
		Title:       sampleMeeting.Title,
		Description: sampleMeeting.Description,
		HasBotGuest: false,
		StartTime:   currentTime,
		EndTime:     currentTime.Add(time.Minute * time.Duration(meetingDuration)),
		CreatedAt:   currentTime,
		Hosts:       hosts,
	}

	// Create Meeting
	createdMeeting, err := repo.CreateMeeting(newMeeting)
	assert.NoError(t, err, "expected no error")
	assert.NotNil(t, createdMeeting, "expected meeting to be created")
	assert.Equal(t, createdMeeting.Duration, sampleMeeting.Duration, "expected duration to match")
	assert.Equal(t, createdMeeting.Title, sampleMeeting.Title, "expected title to match")
	assert.Equal(t, createdMeeting.Description, sampleMeeting.Description, "expected description to match")
	assert.Equal(t, newMeeting, createdMeeting, "expected meeting to equal retrieved meeting")
}
