package meeting

import (
	"fmt"
	"time"

	"github.com/rise8-us/neverl8/model"
	"github.com/rise8-us/neverl8/service/calendar"
)

type MeetingRepositoryInterface interface {
	CreateMeeting(meeting *model.Meetings) (*model.Meetings, error)
	GetAllMeetings() ([]model.Meetings, error)
}

type MeetingService struct {
	meetingRepo     MeetingRepositoryInterface
	calendarService calendar.CalendarServiceInterface
}

func NewMeetingService(meetingRepo MeetingRepositoryInterface, calendarService calendar.CalendarServiceInterface) *MeetingService {
	return &MeetingService{meetingRepo: meetingRepo, calendarService: calendarService}
}

func (s *MeetingService) CreateMeeting(meetings *model.Meetings) (*model.Meetings, error) {
	return s.meetingRepo.CreateMeeting(meetings)
}

func (s *MeetingService) GetAllMeetings() ([]model.Meetings, error) {
	return s.meetingRepo.GetAllMeetings()
}

func (s *MeetingService) GetAvailableTimeBlocks(meeting *model.Meetings, date time.Time) ([]model.TimePreference, error) {
	if date.Day() == time.Now().Day() {
		return []model.TimePreference{}, fmt.Errorf("cannot schedule a meeting less than one day in advance")
	}

	potentialHosts := meeting.Hosts

	// Remove the host that most recently hosted a meeting, assuming there's more than one host
	if len(potentialHosts) > 1 {
		potentialHosts = RemoveMostRecentHost(potentialHosts)
	}

	calendarEvents := s.calendarService.GetAllCalendarEventsForDay(date, potentialHosts)
	eventMap := MapCalendarEventsByHost(calendarEvents)

	availableTimeBlocks := []model.TimePreference{}
	for i := range potentialHosts {
		availableTimeBlocks = append(availableTimeBlocks, potentialHosts[i].TimePreferences...)
	}

	deconflictedTimeBlocks := DetermineAvailableTimeSlots(availableTimeBlocks, eventMap)

	return deconflictedTimeBlocks, nil
}

func RemoveMostRecentHost(potentialHosts []model.Host) []model.Host {
	var mostRecentHost *model.Host
	var mostRecentTime time.Time
	// Find most recent host
	for i := range potentialHosts {
		if potentialHosts[i].LastMeetingTime.After(mostRecentTime) {
			mostRecentTime = potentialHosts[i].LastMeetingTime
			mostRecentHost = &potentialHosts[i]
		}
	}

	// Remove the most recent host
	for i := range potentialHosts {
		if potentialHosts[i].ID == mostRecentHost.ID {
			potentialHosts = append(potentialHosts[:i], potentialHosts[i+1:]...)
			break
		}
	}

	return potentialHosts
}

func MapCalendarEventsByHost(events []model.CalendarEvent) map[uint][]model.CalendarEvent {
	eventMap := make(map[uint][]model.CalendarEvent)
	for _, event := range events {
		eventMap[event.HostID] = append(eventMap[event.HostID], event)
	}
	return eventMap
}

func DetermineAvailableTimeSlots(timeSlots []model.TimePreference, eventMap map[uint][]model.CalendarEvent) []model.TimePreference {
	var availableSlots []model.TimePreference
	layout := "15:04"

	for _, slot := range timeSlots {
		startTime, _ := time.Parse(layout, slot.StartWindow)
		endTime, _ := time.Parse(layout, slot.EndWindow)

		conflicts := false
		for _, event := range eventMap[slot.HostID] {
			if !(startTime.After(event.EndTime) || endTime.Before(event.StartTime)) {
				conflicts = true
				break
			}
		}
		if !conflicts {
			availableSlots = append(availableSlots, slot)
		}
	}

	return availableSlots
}
