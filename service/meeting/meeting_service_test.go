package meeting_test

import (
	"testing"

	"github.com/rise8-us/neverl8/model"
	meetingSvc "github.com/rise8-us/neverl8/service/meeting"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockMeetingRepository struct {
	mock.Mock
}

func (m *MockMeetingRepository) CreateMeeting(meeting *model.Meetings, host []model.Host) (*model.Meetings, error) {
	args := m.Called(meeting, host)
	return args.Get(0).(*model.Meetings), args.Error(1)
}

func (m *MockMeetingRepository) GetAllMeetings() ([]model.Meetings, error) {
	args := m.Called()
	return args.Get(0).([]model.Meetings), args.Error(1)
}

func TestMeetingService_CreateMeeting(t *testing.T) {
	// Create a new instance of our mock repository
	mockRepo := new(MockMeetingRepository)
	meetingService := meetingSvc.NewMeetingService(mockRepo)

	// Setup expectations
	meeting := &model.Meetings{}
	hosts := []model.Host{{}, {}}
	mockRepo.On("CreateMeeting", meeting, hosts).Return(meeting, nil)

	// Call the service method
	result, err := meetingService.CreateMeeting(meeting, hosts)

	// Assert expectations
	assert.NoError(t, err)
	assert.Equal(t, meeting, result)
	mockRepo.AssertExpectations(t)
}

func TestMeetingService_GetAllMeetings(t *testing.T) {
	mockRepo := new(MockMeetingRepository)
	meetingService := meetingSvc.NewMeetingService(mockRepo)

	// Setup expectations
	meetings := []model.Meetings{{}, {}}
	mockRepo.On("GetAllMeetings").Return(meetings, nil)

	// Call the service method
	result, err := meetingService.GetAllMeetings()

	// Assert expectations
	assert.NoError(t, err)
	assert.Equal(t, meetings, result)
	mockRepo.AssertExpectations(t) // Verify that all expectations were met
}
