package tests

import (
	"time"

	"github.com/rise8-us/neverl8/model"
	"github.com/stretchr/testify/mock"
)

type MockMeetingService struct {
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

func (m *MockMeetingService) GetMeetingByID(id uint) (*model.Meetings, error) {
	args := m.Called(id)
	return args.Get(0).(*model.Meetings), args.Error(1)
}

func (m *MockMeetingService) GetMeetingsByDate(date string) ([]model.Meetings, error) {
	args := m.Called(date)
	return args.Get(0).([]model.Meetings), args.Error(1)
}

func (m *MockMeetingService) GetAvailableTimeBlocks(meeting *model.Meetings, day time.Time) ([]model.TimePreference, error) {
	args := m.Called(meeting, day)
	return args.Get(0).([]model.TimePreference), args.Error(1)
}

// func (m *MockCalendarService) GetAllCalendarEventsForDay(day time.Time, hosts []model.Host) ([]model.CalendarEvent, error) {
// 	args := m.Called(day, hosts)
// 	return args.Get(0).([]model.CalendarEvent), args.Error(1)
// }

func (m *MockMeetingService) UpdateMeeting(meeting *model.Meetings) error {
	args := m.Called(meeting)
	return args.Error(0)
}
