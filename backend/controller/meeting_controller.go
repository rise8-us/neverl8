package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/rise8-us/neverl8/model"
	meetingSvc "github.com/rise8-us/neverl8/service/meeting"
)

type MeetingController struct {
	meetingService *meetingSvc.MeetingService
}

type UpdateMeetingTimeRequest struct {
	MeetingID      uint                 `json:"meeting_id"`
	TimeSlot       model.TimePreference `json:"time_slot"`
	CandidateName  string               `json:"candidate_name"`
	CandidateEmail string               `json:"candidate_email"`
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
			_, writeErr := w.Write([]byte(err.Error()))
			if writeErr != nil {
				log.Printf("Error writing response: %v", writeErr)
			}
			return
		}
	} else {
		meetings, err = mc.meetingService.GetAllMeetings()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, writeErr := w.Write([]byte(err.Error()))
			if writeErr != nil {
				log.Printf("Error writing response: %v", writeErr)
			}
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(meetings); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, writeErr := w.Write([]byte(err.Error()))
		if writeErr != nil {
			log.Printf("Error writing response: %v", writeErr)
		}
	}
}

func (mc *MeetingController) GetMeetingByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))

	meeting, err := mc.meetingService.GetMeetingByID(uint(id))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, writeErr := w.Write([]byte(err.Error()))
		if writeErr != nil {
			log.Printf("Error writing response: %v", writeErr)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(meeting); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, writeErr := w.Write([]byte(err.Error()))
		if writeErr != nil {
			log.Printf("Error writing response: %v", writeErr)
		}
	}
}

func (mc *MeetingController) GetAvailableTimeBlocks(w http.ResponseWriter, r *http.Request) {
	date, err := time.Parse("2006-01-02", r.URL.Query().Get("date"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, writeErr := w.Write([]byte(err.Error()))
		if writeErr != nil {
			log.Printf("Error writing response: %v", writeErr)
		}
		return
	}

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, writeErr := w.Write([]byte(err.Error()))
		if writeErr != nil {
			log.Printf("Error writing response: %v", writeErr)
		}
		return
	}

	meeting, err := mc.meetingService.GetMeetingByID(uint(id))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, writeErr := w.Write([]byte(err.Error()))
		if writeErr != nil {
			log.Printf("Error writing response: %v", writeErr)
		}
		return
	}

	timeBlocks, err := mc.meetingService.GetAvailableTimeBlocks(meeting, date)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, writeErr := w.Write([]byte(err.Error()))
		if writeErr != nil {
			log.Printf("Error writing response: %v", writeErr)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(timeBlocks); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, writeErr := w.Write([]byte(err.Error()))
		if writeErr != nil {
			log.Printf("Error writing response: %v", writeErr)
		}
	}
}

func (mc *MeetingController) UpdateMeetingTime(w http.ResponseWriter, r *http.Request) {
	var req UpdateMeetingTimeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, writeErr := w.Write([]byte(err.Error()))
		if writeErr != nil {
			log.Printf("Error writing response: %v", writeErr)
		}
		return
	}

	meeting, err := mc.meetingService.GetMeetingByID(req.MeetingID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, writeErr := w.Write([]byte(err.Error()))
		if writeErr != nil {
			log.Printf("Error writing response: %v", writeErr)
		}
		return
	}

	meeting.StartTime = req.TimeSlot.StartTime
	meeting.EndTime = req.TimeSlot.EndTime

	// Update the meeting with the new times
	// TODO: Update candidate information (name and email) in the database
	err = mc.meetingService.UpdateMeeting(meeting)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, writeErr := w.Write([]byte(err.Error()))
		if writeErr != nil {
			log.Printf("Error writing response: %v", writeErr)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(meeting); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, writeErr := w.Write([]byte(err.Error()))
		if writeErr != nil {
			log.Printf("Error writing response: %v", writeErr)
		}
	}
}
