package cli

import (
	"fmt"
	"log"
	"time"

	"github.com/rise8-us/neverl8/model"
	"github.com/rise8-us/neverl8/service"
)

type CLI struct {
	meetingService *service.MeetingService
}

func NewCLI(meetingService *service.MeetingService) *CLI {
	return &CLI{meetingService}
}

func (c *CLI) CreateMeetingFromCLI() {
	// Create new Hosts
	hosts := []*model.Hosts{
		{HostName: "Host 1"},
		{HostName: "Host 2"},
	}

	// New Meeting to be created
	newMeeting := &model.Meetings{
		CandidateId: 2,
		Calendar:    "Example Calendar",
		Duration:    60,
		Title:       "Example Session",
		Description: "Discuss the future of NeverL8",
		HasBotGuest: false,
		StartTime:   time.Now(),
		EndTime:     time.Now().Add(time.Minute * 60),
	}

	// Create Meeting and Hosts
	createdMeeting, err := c.meetingService.CreateMeeting(newMeeting, hosts)
	if err != nil {
		log.Fatalf("Failed to create meeting and hosts: %v", err)
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
	for i, meeting := range meetings {
		fmt.Printf("%d. %s\n", i+1, meeting.Title)
	}
}

// func (c *CLI) UpdateMeetingFromCLI() {

// }

// func (c * CLI) DeleteMeetingFromCLI() {

// }
