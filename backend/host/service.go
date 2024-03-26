package host

import (
	"fmt"

	"github.com/rise8-us/neverl8/model"
)

type HostRepositoryInterface interface {
	CreateHost(host *model.Host) (*model.Host, error)
	GetHostByID(id uint) (*model.Host, error)
	GetAllHosts() ([]model.Host, error)
	CreateTimePreference(timePreference *model.TimePreference) (*model.TimePreference, error)
	CreateCalendar(calendar *model.Calendar, host *model.Host) (*model.Calendar, error)
}

type HostService struct {
	hostRepo HostRepositoryInterface
}

func NewHostService(hostRepo HostRepositoryInterface) *HostService {
	return &HostService{hostRepo}
}

func (s *HostService) CreateHost(host *model.Host) (*model.Host, error) {
	return s.hostRepo.CreateHost(host)
}

func (s *HostService) GetAllHosts() ([]model.Host, error) {
	return s.hostRepo.GetAllHosts()
}

func (s *HostService) GetHostByID(id uint) (*model.Host, error) {
	return s.hostRepo.GetHostByID(id)
}

func (s *HostService) CreateTimePreference(timePreference *model.TimePreference) (*model.TimePreference, error) {
	if timePreference.StartTime.IsZero() || timePreference.EndTime.IsZero() {
		return timePreference, fmt.Errorf("invalid time preference: start and end times must be provided")
	}
	return s.hostRepo.CreateTimePreference(timePreference)
}

func (s *HostService) CreateCalendar(calendar *model.Calendar, host *model.Host) (*model.Calendar, error) {
	return s.hostRepo.CreateCalendar(calendar, host)
}
