package cli

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

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
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Creating a new meeting...")
	fmt.Print("Enter calendar: ")
	calendar, _ := reader.ReadString('\n')
	fmt.Print("Enter duration (minutes): ")
	durationStr, _ := reader.ReadString('\n')
	duration, _ := strconv.Atoi(strings.TrimSpace(durationStr))
	fmt.Print("Enter title: ")
	title, _ := reader.ReadString('\n')
	fmt.Print("Enter description: ")
	description, _ := reader.ReadString('\n')
	fmt.Print("Enter hosts: ")
	hosts, _ := reader.ReadString('\n')
	fmt.Print("Has bot guest? (true/false): ")
	hasBotGuestStr, _ := reader.ReadString('\n')
	hasBotGuest := strings.TrimSpace(hasBotGuestStr) == "true"

	meeting := &model.Meeting{
		Calendar:    strings.TrimSpace(calendar),
		Duration:    duration,
		Title:       strings.TrimSpace(title),
		Description: strings.TrimSpace(description),
		Hosts:       strings.TrimSpace(hosts),
		HasBotGuest: hasBotGuest,
	}

	if _, err := c.meetingService.CreateMeeting(meeting); err != nil {
		fmt.Println("Failed to create meeting:", err)
		return
	}

	fmt.Println("Meeting created successfully!")
}
