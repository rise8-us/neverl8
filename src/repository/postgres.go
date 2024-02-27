package hello

import (
	"context"

	"github.com/pkg/errors"

	"github.com/drewfugate/neverl8/model"
	"github.com/jinzhu/gorm"
)

var ErrNotExist = errors.New("hello does not exist") // Define the error variable

type PostgresRepo struct {
	DB *gorm.DB
}

func (r *PostgresRepo) GetHello(ctx context.Context) (model.Hello, error) {
	return model.Hello{}, nil
}

type MeetingRepository struct {
	db *gorm.DB
}

func NewMeetingRepository(db *gorm.DB) *MeetingRepository {
	return &MeetingRepository{db}
}

func (r *MeetingRepository) CreateMeeting(meeting *model.Meeting) error {
	if err := r.db.Create(meeting).Error; err != nil {
		return errors.Wrap(err, "failed to create meeting")
	}
	return nil
}

func (r *MeetingRepository) GetMeetingByID(id uint) (*model.Meeting, error) {
	var meeting model.Meeting
	if err := r.db.First(&meeting, id).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, errors.Wrap(err, "failed to get meeting by ID")
	}
	return &meeting, nil
}

func (r *MeetingRepository) UpdateMeeting(meeting *model.Meeting) error {
	if err := r.db.Save(meeting).Error; err != nil {
		return errors.Wrap(err, "failed to update meeting")
	}
	return nil
}

func (r *MeetingRepository) DeleteMeeting(id uint) error {
	if err := r.db.Delete(&model.Meeting{}, id).Error; err != nil {
		return errors.Wrap(err, "failed to delete meeting")
	}
	return nil
}
