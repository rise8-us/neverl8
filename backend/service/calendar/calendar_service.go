package calendar

import (
	"time"

	"github.com/rise8-us/neverl8/model"
)

type CalendarServiceInterface interface {
	GetAllCalendarEventsForDay(day time.Time, hosts []model.Host) ([]model.CalendarEvent, error)
}

// TODO: This function should call the google api to get the calendar for each host
func GetAllCalendarEventsForDay(_ time.Time, _ []model.Host) ([]model.CalendarEvent, error) {
	return nil, nil
}
