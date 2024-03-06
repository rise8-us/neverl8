package service

import (
	"github.com/rise8-us/neverl8/model"
	repository "github.com/rise8-us/neverl8/repository"
)

type MeetingService struct {
	meetingRepo *repository.MeetingRepository
}

func NewMeetingService(meetingRepo *repository.MeetingRepository) *MeetingService {
	return &MeetingService{meetingRepo}
}

func (s *MeetingService) CreateMeeting(meetings *model.Meetings, hosts *[]model.Hosts) (*model.Meetings, error) {
	return s.meetingRepo.CreateMeeting(meetings, *hosts)
}

func (s *MeetingService) GetAllMeetings() ([]model.Meetings, error) {
	return s.meetingRepo.GetAllMeetings()
}
