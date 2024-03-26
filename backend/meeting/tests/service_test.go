package tests_test

import (
	"testing"
	"time"

	"github.com/rise8-us/neverl8/meeting"
	"github.com/rise8-us/neverl8/meeting/tests"
	"github.com/rise8-us/neverl8/model"
	"github.com/stretchr/testify/assert"
)

func TestMeetingService_CreateMeeting(t *testing.T) {
	// Create a new instance of our mock repository
	mockRepo := new(tests.MockMeetingService)
	meetingService := meeting.NewMeetingService(mockRepo, nil)

	// Setup expectations
	meeting := &model.Meetings{}
	mockRepo.On("CreateMeeting", meeting).Return(meeting, nil)

	// Call the service method
	result, err := meetingService.CreateMeeting(meeting)

	// Assert expectations
	assert.NoError(t, err)
	assert.Equal(t, meeting, result)
	mockRepo.AssertExpectations(t)
}

func TestMeetingService_GetAllMeetings(t *testing.T) {
	mockRepo := new(tests.MockMeetingService)
	meetingService := meeting.NewMeetingService(mockRepo, nil)

	// Setup expectations
	meetings := []model.Meetings{{}, {}}
	mockRepo.On("GetAllMeetings").Return(meetings, nil)

	// Call the service method
	result, err := meetingService.GetAllMeetings()

	// Assert expectations
	assert.NoError(t, err)
	assert.Equal(t, meetings, result)
	mockRepo.AssertExpectations(t)
}

func TestMeetingService_GetAvailableTimeBlocks(t *testing.T) {
	mockRepo := new(tests.MockMeetingService)
	meetingService := meeting.NewMeetingService(mockRepo, nil)

	// Create a meeting with a host that has a time preference from 0900 to 1000 and 1200 to 1500 a day in advance
	date, _ := time.Parse("2006-01-02", time.Now().Format("2006-01-02"))

	// Create sample time preferences for the host
	layout := "2006-01-02-15:04"
	startTimeHost1, _ := time.Parse(layout, "9:00")
	endTimeHost1, _ := time.Parse(layout, "17:00")
	startTimeHost2, _ := time.Parse(layout, "10:00")
	endTimeHost2, _ := time.Parse(layout, "18:00")

	hosts := []model.Host{{ID: 1, LastMeetingTime: date.AddDate(0, 0, -3),
		TimePreferences: []model.TimePreference{{HostID: 1, StartTime: startTimeHost1, EndTime: endTimeHost1}}},
		{ID: 2, LastMeetingTime: date.AddDate(0, 0, -1),
			TimePreferences: []model.TimePreference{{HostID: 2, StartTime: startTimeHost2, EndTime: endTimeHost2}}}}
	meeting := model.Meetings{Duration: 60, Hosts: hosts}

	// TODO: Reimplement this mock upon completion of the calendar service's google api implementation
	// mockCalendarService.On("GetAllCalendarEventsForDay", date, mock.Anything).Return(calendarEvents, nil)

	result, err := meetingService.GetAvailableTimeBlocks(&meeting, date.AddDate(0, 0, 1))
	assert.NoError(t, err)

	// Should return available times between 1100 to 1700 in a slice with 60 minute intervals
	projectedStartTime, _ := time.Parse(layout, "11:00")
	projectedEndTime, _ := time.Parse(layout, "12:00")
	for _, result := range result {
		assert.Equal(t, model.TimePreference{HostID: 1, StartTime: projectedStartTime, EndTime: projectedEndTime}, result)
		projectedStartTime = projectedEndTime
		projectedEndTime = projectedEndTime.Add(time.Hour)
	}

	// Should remove the host with the most recent meeting
	for _, host := range result {
		assert.NotEqual(t, uint(2), host.HostID)
	}

	// Test invalid date (less than 1 day in advance)
	date = time.Now()
	result, err = meetingService.GetAvailableTimeBlocks(&meeting, date)
	assert.Error(t, err)
	assert.Empty(t, result)
	assert.Equal(t, err.Error(), "cannot schedule a meeting less than one day in advance")
}
