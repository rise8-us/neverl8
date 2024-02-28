package repository

import (
	"github.com/pkg/errors"

	"github.com/drewfugate/neverl8/model"
	"github.com/jinzhu/gorm"
)

type MeetingRepository struct {
	DB *gorm.DB
}

func NewMeetingRepository(db *gorm.DB) *MeetingRepository {
	return &MeetingRepository{DB: db}
}

func (r *MeetingRepository) CreateMeeting(meeting *model.Meeting) (*model.Meeting, error) {
	if err := r.DB.Create(meeting).Error; err != nil {
		return nil, errors.Wrap(err, "failed to create meeting")
	}
	return meeting, nil
}

func (r *MeetingRepository) GetMeetingByID(id uint) (*model.Meeting, error) {
	var meeting model.Meeting
	if err := r.DB.First(&meeting, id).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, errors.Wrap(err, "failed to get meeting by ID")
	}
	return &meeting, nil
}

func (r *MeetingRepository) UpdateMeeting(meeting *model.Meeting) error {
	if err := r.DB.Save(meeting).Error; err != nil {
		return errors.Wrap(err, "failed to update meeting")
	}
	return nil
}

func (r *MeetingRepository) DeleteMeeting(id uint) error {
	if err := r.DB.Delete(&model.Meeting{}, id).Error; err != nil {
		return errors.Wrap(err, "failed to delete meeting")
	}
	return nil
}
