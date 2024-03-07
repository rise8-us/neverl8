package meeting

import (
	"github.com/rise8-us/neverl8/model"
	repository "github.com/rise8-us/neverl8/repository"
)

type MeetingService struct {
	meetingRepo repository.MeetingRepositoryInterface
}

func NewMeetingService(meetingRepo repository.MeetingRepositoryInterface) *MeetingService {
	return &MeetingService{meetingRepo}
}

func (s *MeetingService) CreateMeeting(meetings *model.Meetings, hosts []model.Host) (*model.Meetings, error) {
	return s.meetingRepo.CreateMeeting(meetings, hosts)
}

func (s *MeetingService) GetAllMeetings() ([]model.Meetings, error) {
	return s.meetingRepo.GetAllMeetings()
}
