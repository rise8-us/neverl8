package meeting_controller

import (
	"github.com/rise8-us/neverl8/service"
)

type MeetingController struct {
	MeetingService *service.MeetingService
}

func NewMeetingController(MeetingService *service.MeetingService) *MeetingController {
	return &MeetingController{MeetingService}
}

//TODO: Add controller methods for HTTP requests
