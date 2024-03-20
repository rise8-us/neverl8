package cli

import (
	"fmt"
	"log"
	"time"

	"github.com/rise8-us/neverl8/model"
	hostSvc "github.com/rise8-us/neverl8/service/host"
	meetingSvc "github.com/rise8-us/neverl8/service/meeting"
)

type CLI struct {
	meetingService *meetingSvc.MeetingService
	hostService    *hostSvc.HostService
}

func NewCLI(meetingService *meetingSvc.MeetingService, hostService *hostSvc.HostService) *CLI {
	return &CLI{meetingService, hostService}
}

func (c *CLI) CreateMeetingFromCLI() {
	// Create new Hosts
	layout := "15:04"
	startTime, _ := time.Parse(layout, "09:00")
	endTime, _ := time.Parse(layout, "17:00")
	startTime2, _ := time.Parse(layout, "11:00")
	endTime2, _ := time.Parse(layout, "15:00")
	hosts := []model.Host{
		{HostName: "Host 1", TimePreferences: []model.TimePreference{{HostID: 1, StartTime: startTime, EndTime: endTime}}},
		{HostName: "Host 2", TimePreferences: []model.TimePreference{{HostID: 2, StartTime: startTime2, EndTime: endTime2}}},
	}
	host1, _ := c.hostService.CreateHost(&hosts[0])
	host2, _ := c.hostService.CreateHost(&hosts[1])

	meetingDuration := 60
	// New Meeting to be created
	newMeeting := &model.Meetings{
		CandidateID: 2,
		Calendar:    "Example Calendar",
		Duration:    60,
		Title:       "Example Session",
		Description: "Discuss the future of NeverL8",
		HasBotGuest: false,
		StartTime:   time.Now(),
		EndTime:     time.Now().Add(time.Minute * time.Duration(meetingDuration)),
		Hosts:       []model.Host{*host1, *host2},
	}

	// Create Meeting and Hosts
	createdMeeting, err := c.meetingService.CreateMeeting(newMeeting)
	if err != nil {
		log.Fatalf("Failed to create meeting: %v", err)
	}

	fmt.Printf("Meeting and hosts created successfully: %+v\n", createdMeeting)
}

func (c *CLI) GetAllMeetingsFromCLI() {
	fmt.Println("The following meetings are on the calendar: ")

	meetings, err := c.meetingService.GetAllMeetings()
	if err != nil {
		fmt.Println("Failed to retrieve all meetings: ", err)
		return
	}

	fmt.Println("Meetings:")
	for i := range meetings {
		fmt.Printf("%d. %s\n", i+1, meetings[i].Title)
	}
}

// func (c *CLI) UpdateMeetingFromCLI() {

// }

// func (c * CLI) DeleteMeetingFromCLI() {

// }
