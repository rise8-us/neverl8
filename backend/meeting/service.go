package meeting

import (
	"fmt"
	"time"

	"github.com/rise8-us/neverl8/calendar"
	"github.com/rise8-us/neverl8/model"
)

type MeetingRepositoryInterface interface {
	CreateMeeting(meeting *model.Meetings) (*model.Meetings, error)
	GetAllMeetings() ([]model.Meetings, error)
	GetMeetingByID(id uint) (*model.Meetings, error)
	GetMeetingsByDate(date string) ([]model.Meetings, error)
	UpdateMeeting(meeting *model.Meetings) error
}

type MeetingService struct {
	meetingRepo MeetingRepositoryInterface
}

func NewMeetingService(meetingRepo MeetingRepositoryInterface, calendarService calendar.CalendarServiceInterface) *MeetingService {
	return &MeetingService{meetingRepo: meetingRepo}
}

func (s *MeetingService) CreateMeeting(meeting *model.Meetings) (*model.Meetings, error) {
	return s.meetingRepo.CreateMeeting(meeting)
}

func (s *MeetingService) GetAllMeetings() ([]model.Meetings, error) {
	return s.meetingRepo.GetAllMeetings()
}

func (s *MeetingService) GetMeetingByID(id uint) (*model.Meetings, error) {
	return s.meetingRepo.GetMeetingByID(id)
}

func (s *MeetingService) GetMeetingsByDate(date string) ([]model.Meetings, error) {
	return s.meetingRepo.GetMeetingsByDate(date)
}

func (s *MeetingService) UpdateMeeting(meeting *model.Meetings) error {
	return s.meetingRepo.UpdateMeeting(meeting)
}

func (s *MeetingService) GetAvailableTimeBlocks(meeting *model.Meetings, date time.Time) ([]model.TimePreference, error) {
	if date.Before(time.Now()) {
		return []model.TimePreference{}, fmt.Errorf("cannot schedule a meeting less than one day in advance")
	}

	potentialHosts := meeting.Hosts

	// Remove the host that most recently hosted a meeting, assuming there's more than one host
	if len(potentialHosts) > 1 {
		potentialHosts = removeMostRecentHost(potentialHosts)
		meeting.Hosts = potentialHosts
		// TODO: remove host from meeting in the database.
		// Also, this check should be moved to CreateMeeting as this function will be called multiple times.
	}

	calendarEvents := []model.CalendarEvent{} // TODO: Call the calendar service to get events for each host
	eventMap := mapCalendarEventsByHost(calendarEvents)

	var availableTimeBlocks []model.TimePreference
	for i := range potentialHosts {
		hostEvents := eventMap[potentialHosts[i].ID]
		hostSlots := generateTimeSlotsForHost(&potentialHosts[i], meeting.Duration, hostEvents)
		for i, slot := range hostSlots {
			// Adjust each slot to use the specified date but keep the original times.
			// The autogenerated 0000-00-00 date from postgres caused issues in the frontend.
			adjustedStart := time.Date(date.Year(), date.Month(), date.Day(), slot.StartTime.Hour(),
				slot.StartTime.Minute(), 0, 0, slot.StartTime.Location())
			adjustedEnd := time.Date(date.Year(), date.Month(), date.Day(), slot.EndTime.Hour(),
				slot.EndTime.Minute(), 0, 0, slot.EndTime.Location())
			adjustedSlot := model.TimePreference{ID: uint(i), HostID: slot.HostID,
				StartTime: adjustedStart, EndTime: adjustedEnd}
			availableTimeBlocks = append(availableTimeBlocks, adjustedSlot)
		}
	}

	return availableTimeBlocks, nil
}

func removeMostRecentHost(potentialHosts []model.Host) []model.Host {
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

func mapCalendarEventsByHost(events []model.CalendarEvent) map[uint][]model.CalendarEvent {
	eventMap := make(map[uint][]model.CalendarEvent)
	for _, event := range events {
		eventMap[event.HostID] = append(eventMap[event.HostID], event)
	}
	return eventMap
}

func generateTimeSlotsForHost(host *model.Host, meetingDuration int, events []model.CalendarEvent) []model.TimePreference {
	if meetingDuration == 0 {
		return []model.TimePreference{}
	}

	var slots []model.TimePreference
	for _, pref := range host.TimePreferences {
		startTime := pref.StartTime
		for {
			// Calculate end time for this slot based on meetingDuration
			endTime := startTime.Add(time.Minute * time.Duration(meetingDuration))
			// Check if this slot goes beyond the preference's end time
			if endTime.After(pref.EndTime) {
				break // This preference's time range is fully processed
			}

			// Check for conflicts with existing events
			conflict := false
			for _, event := range events {
				if endTime.After(event.StartTime) && startTime.Before(event.EndTime) {
					conflict = true
					break
				}
			}

			if !conflict {
				slots = append(slots, model.TimePreference{HostID: host.ID, StartTime: startTime, EndTime: endTime})
			}
			// Move to the next potential slot start time
			startTime = endTime
		}
	}
	return slots
}
