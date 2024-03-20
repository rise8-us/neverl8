package controller

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/rise8-us/neverl8/model"
	meetingSvc "github.com/rise8-us/neverl8/service/meeting"
)

type MeetingController struct {
	meetingService *meetingSvc.MeetingService
}

func NewMeetingController(meetingService *meetingSvc.MeetingService) *MeetingController {
	return &MeetingController{meetingService}
}

func (mc *MeetingController) GetAllMeetings(w http.ResponseWriter, r *http.Request) {
	date := r.URL.Query().Get("date")

	var meetings []model.Meetings
	var err error

	if date != "" {
		meetings, err = mc.meetingService.GetMeetingsByDate(date)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
	} else {
		meetings, err = mc.meetingService.GetAllMeetings()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(meetings); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
}

func (mc *MeetingController) GetMeetingByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))

	meeting, err := mc.meetingService.GetMeetingByID(uint(id))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(meeting); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
}

func (mc *MeetingController) GetAvailableTimeBlocks(w http.ResponseWriter, r *http.Request) {
	date, err := time.Parse("2006-01-02", r.URL.Query().Get("date"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	meeting, err := mc.meetingService.GetMeetingByID(uint(id))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	timeBlocks, err := mc.meetingService.GetAvailableTimeBlocks(*meeting, date)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(timeBlocks); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
}
