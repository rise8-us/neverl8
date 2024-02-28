package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/drewfugate/neverl8/model"
	repository "github.com/drewfugate/neverl8/repository"
	"github.com/go-chi/chi"
)

type MeetingHandler struct {
	meetingRepo *repository.MeetingRepository
}

func NewMeetingHandler(meetingRepo *repository.MeetingRepository) *MeetingHandler {
	return &MeetingHandler{meetingRepo}
}

func (h *MeetingHandler) CreateMeeting(w http.ResponseWriter, r *http.Request) {
	var meeting model.Meeting
	if err := json.NewDecoder(r.Body).Decode(&meeting); err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if _, err := h.meetingRepo.CreateMeeting(&meeting); err != nil {
		http.Error(w, "Failed to create meeting", http.StatusInternalServerError)
		log.Printf("Error creating meeting: %v\n", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(meeting); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		log.Printf("Error encoding response: %v", err)
		return
	}
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

	if err := json.NewEncoder(w).Encode(meeting); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		log.Printf("Error encoding response: %v", err)
		return
	}
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
	if err := json.NewEncoder(w).Encode(meeting); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		log.Printf("Error encoding response: %v", err)
		return
	}
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
