package host_test

import (
	"testing"
	"time"

	"github.com/rise8-us/neverl8/host"
	"github.com/rise8-us/neverl8/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockHostRepository struct {
	mock.Mock
}

func (m *MockHostRepository) CreateHost(host *model.Host) (*model.Host, error) {
	args := m.Called(host)
	return args.Get(0).(*model.Host), args.Error(1)
}

func (m *MockHostRepository) GetAllHosts() ([]model.Host, error) {
	args := m.Called()
	return args.Get(0).([]model.Host), args.Error(1)
}

func (m *MockHostRepository) GetHostByID(id uint) (*model.Host, error) {
	args := m.Called(id)
	return args.Get(0).(*model.Host), args.Error(1)
}

func (m *MockHostRepository) CreateTimePreference(timePreference *model.TimePreference) (*model.TimePreference, error) {
	args := m.Called(timePreference)
	return args.Get(0).(*model.TimePreference), args.Error(1)
}

func (m *MockHostRepository) CreateCalendar(calendar *model.Calendar, host *model.Host) (*model.Calendar, error) {
	args := m.Called(calendar, host)
	return args.Get(0).(*model.Calendar), args.Error(1)
}

func TestHostService_CreateHost(t *testing.T) {
	mockRepo := new(MockHostRepository)
	hostService := host.NewHostService(mockRepo)

	host := &model.Host{}
	mockRepo.On("CreateHost", host).Return(host, nil)

	result, err := hostService.CreateHost(host)

	assert.NoError(t, err, "expected no error")
	assert.Equal(t, host, result, "expected host to be created")
	mockRepo.AssertExpectations(t)
}

func TestHostService_GetAllHosts(t *testing.T) {
	mockRepo := new(MockHostRepository)
	hostService := host.NewHostService(mockRepo)

	hosts := []model.Host{{}, {}}
	mockRepo.On("GetAllHosts").Return(hosts, nil)

	result, err := hostService.GetAllHosts()

	assert.NoError(t, err, "expected no error")
	assert.Equal(t, hosts, result, "expected hosts to be retrieved")
	mockRepo.AssertExpectations(t)
}

func TestHostService_GetHostByID(t *testing.T) {
	mockRepo := new(MockHostRepository)
	hostService := host.NewHostService(mockRepo)

	host := &model.Host{}
	mockRepo.On("GetHostByID", uint(1)).Return(host, nil)

	result, err := hostService.GetHostByID(1)

	assert.NoError(t, err, "expected no error")
	assert.Equal(t, host, result, "expected host to be retrieved")
	mockRepo.AssertExpectations(t)
}

func TestHostService_CreateTimePreference(t *testing.T) {
	mockRepo := new(MockHostRepository)
	hostService := host.NewHostService(mockRepo)

	layout := "15:04"
	startTime, _ := time.Parse(layout, "09:00")
	endTime, _ := time.Parse(layout, "17:00")
	timePreference := model.TimePreference{HostID: 1, StartTime: startTime, EndTime: endTime}
	mockRepo.On("CreateTimePreference", mock.Anything).Return(
		&model.TimePreference{ID: 0, HostID: 1, StartTime: startTime, EndTime: endTime}, nil)

	result, err := hostService.CreateTimePreference(&timePreference)

	assert.NoError(t, err, "expected no error")
	assert.NotNil(t, result, "expected time preferences to be created")
	assert.Equal(t, &timePreference, result, "expected time preferences to be equal")
	mockRepo.AssertExpectations(t)
}

func TestHostService_CreateCalendar(t *testing.T) {
	mockRepo := new(MockHostRepository)
	hostService := host.NewHostService(mockRepo)

	calendar := &model.Calendar{}
	host := &model.Host{}
	mockRepo.On("CreateCalendar", calendar, host).Return(calendar, nil)

	result, err := hostService.CreateCalendar(calendar, host)

	assert.NoError(t, err, "expected no error")
	assert.Equal(t, calendar, result, "expected calendar to be created successfully")
	mockRepo.AssertExpectations(t)
}
