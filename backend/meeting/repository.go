package meeting

import (
	"time"

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
	if err := r.db.Where("id = ?", id).Preload("Hosts.TimePreferences").First(&meeting, id).Error; err != nil {
		return nil, err
	}
	return &meeting, nil
}

func (r *MeetingRepository) GetMeetingsByDate(date string) ([]model.Meetings, error) {
	var meetings []model.Meetings
	// Parse the date query parameter to time.Time
	parsedDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		return nil, err
	}

	dayStart := time.Date(parsedDate.Year(), parsedDate.Month(), parsedDate.Day(), 0, 0, 0, 0, parsedDate.Location())
	dayEnd := dayStart.AddDate(0, 0, 1)

	result := r.db.Where("start_time >= ? AND start_time < ?", dayStart, dayEnd).Preload("Hosts.TimePreferences").Find(&meetings)
	if result.Error != nil {
		return nil, result.Error
	}

	return meetings, nil
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

func (r *MeetingRepository) UpdateMeeting(meeting *model.Meetings) error {
	if err := r.db.Save(meeting).Error; err != nil {
		return err
	}
	return nil
}
