package repository

import (
	"github.com/rise8-us/neverl8/model"
	"gorm.io/gorm"
)

// MeetingRepository handles CRUD operations for Meetings.
type MeetingRepository struct {
	db *gorm.DB
}

// NewMeetingRepository creates a new instance of MeetingRepository.
func NewMeetingRepository(db *gorm.DB) *MeetingRepository {
	return &MeetingRepository{db: db}
}

// Adds a new Meeting to the database, including creating new Host if necessary.
func (r *MeetingRepository) CreateMeeting(meeting *model.Meetings) (*model.Meetings, error) {
	if err := r.db.Create(meeting).Error; err != nil {
		return nil, err
	}

	return meeting, nil
}

// Returns all Meetings from the database.
func (r *MeetingRepository) GetAllMeetings() ([]model.Meetings, error) {
	var meetings []model.Meetings
	if err := r.db.Preload("Hosts").Find(&meetings).Error; err != nil {
		return nil, err
	}
	return meetings, nil
}

// Returns a Meeting by its ID.
func (r *MeetingRepository) GetMeetingByID(id uint) (*model.Meetings, error) {
	var meeting model.Meetings
	if err := r.db.Preload("Hosts").First(&meeting, id).Error; err != nil {
		return nil, err
	}
	return &meeting, nil
}

func (r *MeetingRepository) CreateSampleMeeting(sampleMeeting *model.SampleMeetings) *model.SampleMeetings {
	err := r.db.Create(sampleMeeting)
	if err.Error != nil {
		return nil
	}

	return sampleMeeting
}

func (r *MeetingRepository) GetSampleMeetings() ([]model.SampleMeetings, error) {
	var sampleMeetings []model.SampleMeetings
	if err := r.db.Find(&sampleMeetings).Error; err != nil {
		return nil, err
	}
	return sampleMeetings, nil
}
