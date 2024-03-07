package repository_test

import (
	"testing"
	"time"

	"github.com/rise8-us/neverl8/model"
	"github.com/rise8-us/neverl8/repository"
	"github.com/stretchr/testify/assert"
)

func GetSampleMeeting() (*model.Meetings, *[]model.Host) {
	// Create new Hosts
	hosts := &[]model.Host{
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
		StartTime:   time.Now().UTC(),
		EndTime:     time.Now().UTC().Add(time.Minute * time.Duration(meetingDuration)),
		CreatedAt:   time.Now().UTC(),
	}

	return meeting, hosts
}

func TestCreateMeeting(t *testing.T) {
	t.Cleanup(func() {
		db.Where("1 = 1").Delete(&model.Meetings{})
		db.Where("1 = 1").Delete(&model.Host{})
	})

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
	t.Cleanup(func() {
		db.Where("1 = 1").Delete(&model.Meetings{})
		db.Where("1 = 1").Delete(&model.Host{})
	})

	repo := repository.NewMeetingRepository(db)

	meeting, hosts := GetSampleMeeting()

	// Create Meeting
	_, err := repo.CreateMeeting(meeting, *hosts)
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

	repo := repository.NewMeetingRepository(db)

	meeting, hosts := GetSampleMeeting()

	// Create Meeting
	createdMeeting, err := repo.CreateMeeting(meeting, *hosts)
	assert.NoError(t, err, "expected no error")

	// Get Meeting 1
	retrievedMeeting, err := repo.GetMeetingByID(createdMeeting.ID)
	assert.NoError(t, err, "expected no error")
	assert.NotNil(t, retrievedMeeting, "expected meeting to be retrieved")
	assert.Equal(t, *createdMeeting, *retrievedMeeting, "expected meeting to equal retrieved meeting")
}
