package meeting_test

import (
	"testing"
	"time"

	"github.com/rise8-us/neverl8/model"
	meetingSvc "github.com/rise8-us/neverl8/service/meeting"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockMeetingService struct {
	mock.Mock
}

type MockCalendarService struct {
	mock.Mock
}

func (m *MockMeetingService) CreateMeeting(meeting *model.Meetings) (*model.Meetings, error) {
	args := m.Called(meeting)
	return args.Get(0).(*model.Meetings), args.Error(1)
}

func (m *MockMeetingService) GetAllMeetings() ([]model.Meetings, error) {
	args := m.Called()
	return args.Get(0).([]model.Meetings), args.Error(1)
}

func (m *MockMeetingService) GetAvailableTimeBlocks(meeting *model.Meetings, day time.Time) ([]model.TimePreference, error) {
	args := m.Called(meeting, day)
	return args.Get(0).([]model.TimePreference), args.Error(1)
}

func (m *MockCalendarService) GetAllCalendarEventsForDay(day time.Time, hosts []model.Host) []model.CalendarEvent {
	args := m.Called(day, hosts)
	return args.Get(0).([]model.CalendarEvent)
}

func TestMeetingService_CreateMeeting(t *testing.T) {
	// Create a new instance of our mock repository
	mockRepo := new(MockMeetingService)
	meetingService := meetingSvc.NewMeetingService(mockRepo, nil)

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
	mockRepo := new(MockMeetingService)
	meetingService := meetingSvc.NewMeetingService(mockRepo, nil)

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
	mockCalendarService := new(MockCalendarService)
	meetingService := meetingSvc.NewMeetingService(nil, mockCalendarService)

	// Create a calendar event for the host from 0900 to 1100
	layout := "15:04"
	startTime, _ := time.Parse(layout, "09:00")
	endTime, _ := time.Parse(layout, "11:00")
	calendarEvents := []model.CalendarEvent{{HostID: 1, StartTime: startTime, EndTime: endTime}}

	// Create a meeting with a host that has a time preference from 0900 to 1000 and 1200 to 1500 a day in advance
	date := time.Date(2021, time.January, 1, 0, 0, 0, 0, time.UTC)
	hosts := []model.Host{{ID: 1, LastMeetingTime: date.AddDate(0, 0, -3),
		TimePreferences: []model.TimePreference{{HostID: 1, StartWindow: "09:00", EndWindow: "10:00"},
			{HostID: 1, StartWindow: "12:00", EndWindow: "15:00"}}},
		{ID: 2, LastMeetingTime: date.AddDate(0, 0, -1),
			TimePreferences: []model.TimePreference{{HostID: 2, StartWindow: "11:00", EndWindow: "15:00"}}}}
	meeting := &model.Meetings{Hosts: hosts}

	mockCalendarService.On("GetAllCalendarEventsForDay", date, mock.Anything).Return(calendarEvents, nil)

	result, err := meetingService.GetAvailableTimeBlocks(meeting, date)

	// Should return available times between 1200 to 1500 due to conflict with calendar event at 0900 to 1000
	assert.NoError(t, err)
	assert.Equal(t, hosts[0].TimePreferences[1], result[0])

	// Should remove the host with the most recent meeting
	for _, host := range result {
		assert.NotEqual(t, uint(2), host.HostID)
	}

	// Test invalid date (less than 1 day in advance)
	date = time.Now()
	result, err = meetingService.GetAvailableTimeBlocks(meeting, date)
	assert.Error(t, err)
	assert.Empty(t, result)
	assert.Equal(t, err.Error(), "cannot schedule a meeting less than one day in advance")

	mockCalendarService.AssertExpectations(t)
}