package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/drewfugate/neverl8/model"
	hello "github.com/drewfugate/neverl8/repository"
	"github.com/go-chi/chi"
)

type Hello struct {
	Repo *hello.PostgresRepo
}

type MeetingHandler struct {
	meetingRepo *hello.MeetingRepository
}

func (h *Hello) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	_, err := w.Write([]byte("Hello, World!"))
	if err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}
}

func NewMeetingHandler(meetingRepo *hello.MeetingRepository) *MeetingHandler {
	return &MeetingHandler{meetingRepo}
}

func (h *MeetingHandler) CreateMeeting(w http.ResponseWriter, r *http.Request) {
	var meeting model.Meeting
	if err := json.NewDecoder(r.Body).Decode(&meeting); err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if err := h.meetingRepo.CreateMeeting(&meeting); err != nil {
		http.Error(w, "Failed to create meeting", http.StatusInternalServerError)
		log.Printf("Error creating meeting: %v\n", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(meeting)
}

func (h *MeetingHandler) GetMeeting(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	meetingID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid meeting ID", http.StatusBadRequest)
		return
	}

	meeting, err := h.meetingRepo.GetMeetingByID(uint(meetingID))
	if err != nil {
		http.Error(w, "Failed to get meeting", http.StatusInternalServerError)
		log.Printf("Error getting meeting: %v\n", err)
		return
	}
	if meeting == nil {
		http.Error(w, "Meeting not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(meeting)
}

func (h *MeetingHandler) UpdateMeeting(w http.ResponseWriter, r *http.Request) {
	var meeting model.Meeting
	if err := json.NewDecoder(r.Body).Decode(&meeting); err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if err := h.meetingRepo.UpdateMeeting(&meeting); err != nil {
		http.Error(w, "Failed to update meeting", http.StatusInternalServerError)
		log.Printf("Error updating meeting: %v\n", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(meeting)
}

func (h *MeetingHandler) DeleteMeeting(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	meetingID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid meeting ID", http.StatusBadRequest)
		return
	}

	if err := h.meetingRepo.DeleteMeeting(uint(meetingID)); err != nil {
		http.Error(w, "Failed to delete meeting", http.StatusInternalServerError)
		log.Printf("Error deleting meeting: %v\n", err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
