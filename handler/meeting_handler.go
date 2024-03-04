package meeting_handler

import (
	"github.com/rise8-us/neverl8/service"
)

type MeetingHandler struct {
	MeetingService *service.MeetingService
}

func NewMeetingHandler(MeetingService *service.MeetingService) *MeetingHandler {
	return &MeetingHandler{MeetingService}
}

//TODO: Add handler methods for HTTP requests
