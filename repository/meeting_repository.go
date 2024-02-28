package repository

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/rise8-us/neverl8/model"
)

type MeetingRepository struct {
	DB *gorm.DB
}

func NewMeetingRepository(db *gorm.DB) *MeetingRepository {
	return &MeetingRepository{DB: db}
}

func (r *MeetingRepository) CreateMeeting(meeting *model.Meeting) (*model.Meeting, error) {
	if err := r.DB.Create(meeting).Error; err != nil {
		return nil, fmt.Errorf("failed to create meeting: %w", err)
	}
	return meeting, nil
}

func (r *MeetingRepository) GetMeetingByID(id uint) (*model.Meeting, error) {
	var meeting model.Meeting
	if err := r.DB.First(&meeting, id).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get meeting by ID: %w", err)
	}
	return &meeting, nil
}

func (r *MeetingRepository) UpdateMeeting(meeting *model.Meeting) error {
	if err := r.DB.Save(meeting).Error; err != nil {
		return fmt.Errorf("failed to update meeting: %w", err)
	}
	return nil
}

func (r *MeetingRepository) DeleteMeeting(id uint) error {
	if err := r.DB.Delete(&model.Meeting{}, id).Error; err != nil {
		return fmt.Errorf("failed to delete meeting: %w", err)
	}
	return nil
}
