package service

import (
	"github.com/drewfugate/neverl8/model"
	hello "github.com/drewfugate/neverl8/repository"
)

type MeetingService struct {
	meetingRepo *hello.MeetingRepository
}

func NewMeetingService(meetingRepo *hello.MeetingRepository) *MeetingService {
	return &MeetingService{meetingRepo}
}

func (s *MeetingService) CreateMeeting(meeting *model.Meeting) error {
	return s.meetingRepo.CreateMeeting(meeting)
}

func (s *MeetingService) GetMeetingByID(id uint) (*model.Meeting, error) {
	return s.meetingRepo.GetMeetingByID(id)
}

func (s *MeetingService) UpdateMeeting(meeting *model.Meeting) error {
	return s.meetingRepo.UpdateMeeting(meeting)
}

func (s *MeetingService) DeleteMeeting(id uint) error {
	return s.meetingRepo.DeleteMeeting(id)
}
