package tests_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-chi/chi"
	"github.com/rise8-us/neverl8/meeting"
	"github.com/rise8-us/neverl8/meeting/tests"
	"github.com/rise8-us/neverl8/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var mockService = new(tests.MockMeetingService)
var mc = meeting.NewMeetingController(mockService)

func Test_Create_Meeting(t *testing.T) {
	body, err := json.Marshal(model.Meetings{ID: 1})
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/api/meeting", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}
	responseRecorder := httptest.NewRecorder()
	router := chi.NewRouter()
	router.Mount("/api", mc.RegisterRoutes())

	mockService.On("CreateMeeting", &model.Meetings{ID: 1}).Return(&model.Meetings{ID: 1}, nil)
	router.ServeHTTP(responseRecorder, req)

	actualBody, err := io.ReadAll(responseRecorder.Body)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.Contains(t, string(actualBody), string(body))
	mockService.AssertExpectations(t)
}

func Test_Get_All_Meetings(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/meetings", nil)
	if err != nil {
		t.Fatal(err)
	}
	responseRecorder := httptest.NewRecorder()
	router := chi.NewRouter()
	router.Mount("/api", mc.RegisterRoutes())

	meetings := []model.Meetings{{ID: 1}, {ID: 2}}
	mockService.On("GetAllMeetings").Return(meetings, nil)
	router.ServeHTTP(responseRecorder, req)

	actualBody, err := io.ReadAll(responseRecorder.Body)
	if err != nil {
		t.Fatal(err)
	}

	var actualMeetings []model.Meetings
	err = json.Unmarshal(actualBody, &actualMeetings)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.Contains(t, string(actualBody), "1")
	assert.Contains(t, string(actualBody), "2")
	assert.Len(t, actualMeetings, 2)
	assert.Equal(t, meetings[0].ID, actualMeetings[0].ID)
	assert.Equal(t, meetings[1].ID, actualMeetings[1].ID)
	mockService.AssertExpectations(t)
}

func Test_Get_Meeting_By_ID(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/meeting/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	responseRecorder := httptest.NewRecorder()
	router := chi.NewRouter()
	router.Mount("/api", mc.RegisterRoutes())

	mockService.On("GetMeetingByID", uint(1)).Return(&model.Meetings{ID: 1}, nil)
	router.ServeHTTP(responseRecorder, req)

	actualBody, err := io.ReadAll(responseRecorder.Body)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.Contains(t, string(actualBody), "1")
	mockService.AssertExpectations(t)
}

func Test_Get_Available_Time_Blocks(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/meeting/time-slots/1?date=2023-01-02", nil)
	if err != nil {
		t.Fatal(err)
	}

	responseRecorder := httptest.NewRecorder()
	router := chi.NewRouter()
	router.Mount("/api", mc.RegisterRoutes())

	date, _ := time.Parse("2006-01-02", "2023-01-02")
	startTime, _ := time.Parse("15:04", "09:00")
	endTime, _ := time.Parse("15:04", "12:00")
	startBlock1, _ := time.Parse("15:04", "09:00")
	endBlock1, _ := time.Parse("15:04", "10:00")
	startBlock2, _ := time.Parse("15:04", "10:00")
	endBlock2, _ := time.Parse("15:04", "11:00")
	startBlock3, _ := time.Parse("15:04", "11:00")
	endBlock3, _ := time.Parse("15:04", "12:00")

	timePreferences := []model.TimePreference{
		{StartTime: startTime, EndTime: endTime},
	}

	mockMeeting := &model.Meetings{
		ID:       1,
		Duration: 60,
		Hosts:    []model.Host{{TimePreferences: timePreferences}},
	}

	mockService.On("GetMeetingByID", uint(1)).Return(mockMeeting, nil)

	expectedTimeBlocks := []model.TimePreference{
		{StartTime: startBlock1, EndTime: endBlock1},
		{StartTime: startBlock2, EndTime: endBlock2},
		{StartTime: startBlock3, EndTime: endBlock3},
	}

	mockService.On("GetAvailableTimeBlocks", mock.Anything, date).Return(expectedTimeBlocks, nil)
	router.ServeHTTP(responseRecorder, req)
	actualBody, err := io.ReadAll(responseRecorder.Body)
	if err != nil {
		t.Fatal(err)
	}

	var actualTimeBlocks []model.TimePreference
	err = json.Unmarshal(actualBody, &actualTimeBlocks)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.Len(t, actualTimeBlocks, len(expectedTimeBlocks))
	for i, expectedBlock := range expectedTimeBlocks {
		assert.Equal(t, expectedBlock.StartTime.Format("15:04"), actualTimeBlocks[i].StartTime.Format("15:04"))
		assert.Equal(t, expectedBlock.EndTime.Format("15:04"), actualTimeBlocks[i].EndTime.Format("15:04"))
	}
	mockService.AssertExpectations(t)
}

func Test_Update_Meeting_Time(t *testing.T) {
	updateReq := meeting.UpdateMeetingTimeRequest{
		MeetingID:      1,
		TimeSlot:       model.TimePreference{StartTime: time.Now(), EndTime: time.Now().Add(time.Hour)},
		CandidateName:  "John Doe",
		CandidateEmail: "johndoe@example.com",
	}

	body, err := json.Marshal(updateReq)
	if err != nil {
		t.Fatal(err)
	}
	req, err := http.NewRequest("PUT", "/api/meeting/schedule", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	responseRecorder := httptest.NewRecorder()
	router := chi.NewRouter()
	router.Mount("/api", mc.RegisterRoutes())

	mockService.On("UpdateMeeting", mock.AnythingOfType("*model.Meetings")).Return(nil)
	mockService.On("GetMeetingByID", uint(1)).Return(&model.Meetings{ID: 1}, nil)

	router.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)

	var updatedMeeting model.Meetings
	err = json.Unmarshal(responseRecorder.Body.Bytes(), &updatedMeeting)
	if err != nil {
		t.Fatal("Failed to unmarshal response body")
	}

	assert.True(t, updatedMeeting.StartTime.Equal(updateReq.TimeSlot.StartTime))
	assert.True(t, updatedMeeting.EndTime.Equal(updateReq.TimeSlot.EndTime))
	mockService.AssertExpectations(t)
}
