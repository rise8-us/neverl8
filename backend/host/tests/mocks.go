package tests

import (
	"github.com/rise8-us/neverl8/model"
	"github.com/stretchr/testify/mock"
)

type MockHostService struct {
	mock.Mock
}

func (m *MockHostService) CreateHost(host *model.Host) (*model.Host, error) {
	args := m.Called(host)
	return args.Get(0).(*model.Host), args.Error(1)
}

func (m *MockHostService) GetAllHosts() ([]model.Host, error) {
	args := m.Called()
	return args.Get(0).([]model.Host), args.Error(1)
}

func (m *MockHostService) GetHostByID(id uint) (*model.Host, error) {
	args := m.Called(id)
	return args.Get(0).(*model.Host), args.Error(1)
}

func (m *MockHostService) CreateTimePreference(timePreference *model.TimePreference) (*model.TimePreference, error) {
	args := m.Called(timePreference)
	return args.Get(0).(*model.TimePreference), args.Error(1)
}

func (m *MockHostService) CreateCalendar(calendar *model.Calendar, host *model.Host) (*model.Calendar, error) {
	args := m.Called(calendar, host)
	return args.Get(0).(*model.Calendar), args.Error(1)
}
