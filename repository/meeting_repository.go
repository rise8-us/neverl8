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

// Adds a new Meeting to the database, including creating new Hosts if necessary.
func (r *MeetingRepository) CreateMeeting(meeting *model.Meetings, hosts []model.Hosts) (*model.Meetings, error) {
	// Step 1: Create and save each Host
	for i := range hosts {
		if err := r.db.Create(&hosts[i]).Error; err != nil {
			return nil, err
		}
	}

	// Step 2: Create the Meeting
	if err := r.db.Create(meeting).Error; err != nil {
		return nil, err
	}

	return meeting, nil
}

// Returns all Meetings from the database.
func (r *MeetingRepository) GetAllMeetings() ([]model.Meetings, error) {
	var meetings []model.Meetings
	if err := r.db.Preload("Hosts").Preload("Hosts.TimePreferences").Find(&meetings).Error; err != nil {
		return nil, err
	}
	return meetings, nil
}
