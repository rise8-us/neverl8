package tests_test

import (
	"testing"
	"time"

	"github.com/rise8-us/neverl8/host"
	"github.com/rise8-us/neverl8/host/tests"
	"github.com/rise8-us/neverl8/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestHostService_CreateHost(t *testing.T) {
	mockRepo := new(tests.MockHostService)
	hostService := host.NewHostService(mockRepo)

	sampleHost := &model.Host{}
	mockRepo.On("CreateHost", sampleHost).Return(sampleHost, nil)

	result, err := hostService.CreateHost(sampleHost)

	assert.NoError(t, err, "expected no error")
	assert.Equal(t, sampleHost, result, "expected host to be created")
	mockRepo.AssertExpectations(t)
}

func TestHostService_GetAllHosts(t *testing.T) {
	mockRepo := new(tests.MockHostService)
	hostService := host.NewHostService(mockRepo)

	hosts := []model.Host{{}, {}}
	mockRepo.On("GetAllHosts").Return(hosts, nil)

	result, err := hostService.GetAllHosts()

	assert.NoError(t, err, "expected no error")
	assert.Equal(t, hosts, result, "expected hosts to be retrieved")
	mockRepo.AssertExpectations(t)
}

func TestHostService_GetHostByID(t *testing.T) {
	mockRepo := new(tests.MockHostService)
	hostService := host.NewHostService(mockRepo)

	sampleHost := &model.Host{}
	mockRepo.On("GetHostByID", uint(1)).Return(sampleHost, nil)

	result, err := hostService.GetHostByID(1)

	assert.NoError(t, err, "expected no error")
	assert.Equal(t, sampleHost, result, "expected host to be retrieved")
	mockRepo.AssertExpectations(t)
}

func TestHostService_CreateTimePreference(t *testing.T) {
	mockRepo := new(tests.MockHostService)
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
	mockRepo := new(tests.MockHostService)
	hostService := host.NewHostService(mockRepo)

	calendar := &model.Calendar{}
	sampleHost := &model.Host{}
	mockRepo.On("CreateCalendar", calendar, sampleHost).Return(calendar, nil)

	result, err := hostService.CreateCalendar(calendar, sampleHost)

	assert.NoError(t, err, "expected no error")
	assert.Equal(t, calendar, result, "expected calendar to be created successfully")
	mockRepo.AssertExpectations(t)
}
