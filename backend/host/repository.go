package host

import (
	"github.com/rise8-us/neverl8/model"
	"gorm.io/gorm"
)

// HostRepository handles CRUD operations for hosts and their time preferences.
type HostRepository struct {
	db *gorm.DB
}

// NewHostRepository creates a new instance of HostRepository.
func NewHostRepository(db *gorm.DB) *HostRepository {
	return &HostRepository{db: db}
}

func (r *HostRepository) CreateHost(host *model.Host) (*model.Host, error) {
	if err := r.db.Create(host).Error; err != nil {
		return nil, err
	}

	return host, nil
}

func (r *HostRepository) GetHostByID(id uint) (*model.Host, error) {
	var host model.Host
	if err := r.db.Preload("TimePreferences").Preload("Calendars").Preload("Meetings").First(&host, id).Error; err != nil {
		return nil, err
	}
	return &host, nil
}

func (r *HostRepository) GetAllHosts() ([]model.Host, error) {
	var host []model.Host
	if err := r.db.Preload("TimePreferences").Preload("Calendars").Preload("Meetings").Find(&host).Error; err != nil {
		return nil, err
	}
	return host, nil
}

// Adds a new TimePreference to the database for a Host.
func (r *HostRepository) CreateTimePreference(timePreference *model.TimePreference) (*model.TimePreference, error) {
	if err := r.db.Create(timePreference).Error; err != nil {
		return timePreference, err
	}

	return timePreference, nil
}

func (r *HostRepository) CreateCalendar(calendar *model.Calendar, host *model.Host) (*model.Calendar, error) {
	if err := r.db.Create(calendar).Error; err != nil {
		return nil, err
	}

	// Create the association between Host and Calendar
	err := r.db.Model(host).Association("Calendars").Append(&calendar)
	if err != nil {
		return nil, err
	}

	return calendar, nil
}
