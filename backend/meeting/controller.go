package meeting

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	"github.com/rise8-us/neverl8/model"
)

type MeetingServiceInterface interface {
	CreateMeeting(meeting *model.Meetings) (*model.Meetings, error)
	GetAllMeetings() ([]model.Meetings, error)
	GetMeetingByID(id uint) (*model.Meetings, error)
	GetMeetingsByDate(date string) ([]model.Meetings, error)
	GetAvailableTimeBlocks(meeting *model.Meetings, date time.Time) ([]model.TimePreference, error)
	UpdateMeeting(meeting *model.Meetings) error
}

type MeetingController struct {
	meetingService MeetingServiceInterface
}

type UpdateMeetingTimeRequest struct {
	MeetingID      uint                 `json:"meeting_id"`
	TimeSlot       model.TimePreference `json:"time_slot"`
	CandidateName  string               `json:"candidate_name"`
	CandidateEmail string               `json:"candidate_email"`
}

func NewMeetingController(meetingService MeetingServiceInterface) *MeetingController {
	return &MeetingController{meetingService: meetingService}
}

func (mc *MeetingController) RegisterRoutes() chi.Router {
	router := chi.NewRouter()

	router.Get("/meetings", mc.getAllMeetings)
	router.Get("/meeting/{id}", mc.getMeetingByID)
	router.Get("/meeting/time-slots/{id}", mc.getAvailableTimeBlocks)
	router.Post("/meeting", mc.createMeeting)
	router.Put("/meeting/schedule", mc.updateMeetingTime)

	return router
}

func (mc *MeetingController) createMeeting(w http.ResponseWriter, r *http.Request) {
	var meeting model.Meetings
	if err := json.NewDecoder(r.Body).Decode(&meeting); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, writeErr := w.Write([]byte(err.Error()))
		if writeErr != nil {
			log.Printf("Error writing response: %v", writeErr)
		}
		return
	}

	_, err := mc.meetingService.CreateMeeting(&meeting)
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

func (mc *MeetingController) getAllMeetings(w http.ResponseWriter, r *http.Request) {
	date := chi.URLParam(r, "date")

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

func (mc *MeetingController) getMeetingByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

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

func (mc *MeetingController) getAvailableTimeBlocks(w http.ResponseWriter, r *http.Request) {
	date, err := time.Parse("2006-01-02", r.URL.Query().Get("date"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, writeErr := w.Write([]byte(err.Error()))
		if writeErr != nil {
			log.Printf("Error writing response: %v", writeErr)
		}
		return
	}

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
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

func (mc *MeetingController) updateMeetingTime(w http.ResponseWriter, r *http.Request) {
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
