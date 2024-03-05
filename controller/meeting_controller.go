package meetingcontroller

import (
	"github.com/rise8-us/neverl8/service"
)

type meetingController struct {
	meetingService *service.MeetingService
}

func NewMeetingController(meetingService *service.MeetingService) *meetingController {
	return &meetingController{meetingService}
}

//TODO: Add controller methods for HTTP requests
