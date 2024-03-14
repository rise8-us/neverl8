package host

import (
	"fmt"
	"time"

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
	// Store time preferences as 15 minute blocks
	timeBlocks := []model.TimePreference{}
	if host.TimePreferences != nil {
		for i := range host.TimePreferences {
			timeBlocks = append(timeBlocks, ConvertTimePreferenceToBlocks(&host.TimePreferences[i])...)
		}
	}
	host.TimePreferences = timeBlocks
	return s.hostRepo.CreateHost(host)
}

func (s *HostService) GetAllHosts() ([]model.Host, error) {
	return s.hostRepo.GetAllHosts()
}

func (s *HostService) GetHostByID(id uint) (*model.Host, error) {
	return s.hostRepo.GetHostByID(id)
}

func (s *HostService) CreateTimePreference(timePreference *model.TimePreference) ([]model.TimePreference, error) {
	// Store time preferences as 15 minute blocks
	if timePreference.StartWindow == "" || timePreference.EndWindow == "" {
		return nil, fmt.Errorf("invalid time preferences")
	}
	timeBlocks := ConvertTimePreferenceToBlocks(timePreference)
	return s.hostRepo.CreateTimePreference(timeBlocks)
}

func (s *HostService) CreateCalendar(calendar *model.Calendar, host *model.Host) (*model.Calendar, error) {
	return s.hostRepo.CreateCalendar(calendar, host)
}

func ConvertTimePreferenceToBlocks(timePreference *model.TimePreference) []model.TimePreference {
	layout := "15:04"
	blockDuration := 15
	startTime, err := time.Parse(layout, timePreference.StartWindow)
	if err != nil {
		return nil
	}
	endTime, err := time.Parse(layout, timePreference.EndWindow)
	if err != nil {
		return nil
	}

	var timeBlocks []model.TimePreference
	for currentTime := startTime; currentTime.Before(endTime); currentTime = currentTime.Add(time.Minute * time.Duration(blockDuration)) {
		blockEndTime := currentTime.Add(time.Minute * time.Duration(blockDuration))
		if blockEndTime.After(endTime) {
			break
		}

		timeBlocks = append(timeBlocks, model.TimePreference{
			HostID:      timePreference.HostID,
			StartWindow: currentTime.Format(layout),
			EndWindow:   blockEndTime.Format(layout),
		})
	}

	return timeBlocks
}
