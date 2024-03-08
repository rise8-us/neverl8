package meetingcontroller

import (
	meetingSvc "github.com/rise8-us/neverl8/service/meeting"
)

type meetingController struct {
	meetingService *meetingSvc.MeetingService
}

func NewMeetingController(meetingService *meetingSvc.MeetingService) *meetingController {
	return &meetingController{meetingService}
}

//TODO: Add controller methods for HTTP requests
