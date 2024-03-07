package host

import (
	"github.com/rise8-us/neverl8/model"
	repository "github.com/rise8-us/neverl8/repository"
)

type HostService struct {
	hostRepo repository.HostRepositoryInterface
}

func NewHostService(hostRepo repository.HostRepositoryInterface) *HostService {
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
	return s.hostRepo.CreateTimePreference(timePreference)
}

func (s *HostService) CreateCalendar(calendar *model.Calendar, host *model.Host) (*model.Calendar, error) {
	return s.hostRepo.CreateCalendar(calendar, host)
}
