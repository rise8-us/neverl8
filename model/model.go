package model

import (
	"time"
)

type Meetings struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	CandidateID uint      `json:"candidate_id" gorm:"type:integer; not null"`
	Calendar    string    `json:"calendar" gorm:"type:varchar(255); not null"`
	Duration    int       `json:"duration" gorm:"type:integer; not null"`
	Title       string    `json:"title" gorm:"type:varchar(255); not null"`
	Description string    `json:"description" gorm:"type:varchar(255); default: no description provided"`
	HasBotGuest bool      `json:"has_bot_guest" gorm:"type:bool; default: false"`
	StartTime   time.Time `json:"start_time" gorm:"type: timestamp; not null"`
	EndTime     time.Time `json:"end_time" gorm:"type: timestamp; not null"`
	CreatedAt   time.Time `json:"created_at" gorm:"type: timestamp; default: current_timestamp"`
	Hosts       []Host    `gorm:"many2many:host_meetings;joinForeignKey:meeting_id;inverseJoinForeignKey:host_id"`
}

type Host struct {
	ID              uint             `json:"id" gorm:"primaryKey"`
	HostName        string           `json:"host_name" gorm:"type:varchar(255); not null"`
	Candidates      []Candidates     `json:"candidates_id" gorm:"foreignKey:host_id"`
	Meetings        []Meetings       `gorm:"many2many:host_meetings;joinForeignKey:host_id;inverseJoinForeignKey:meeting_id"`
	TimePreferences []TimePreference `json:"time_preferences" gorm:"foreignKey:host_id"` // One to many relationship with time preference
	Calendars       []Calendar       `gorm:"many2many:host_calendars;joinForeignKey:calendar_id;inverseJoinForeignKey:host_id"`
}

type Candidates struct {
	ID              uint   `json:"id" gorm:"primaryKey"`
	HostID          uint   `json:"host_id"` // foreign key to hosts
	CandidateName   string `json:"candidate_name" gorm:"type:varchar(255); not null"`
	Role            string `json:"role" gorm:"type:varchar(255); default: unknown role"`
	Email           string `json:"email" gorm:"type:varchar(255); default: unknown email"`
	PhoneNumber     string `json:"phone_number" gorm:"type:varchar(255); default: unknown phone number"`
	InterviewStatus string `json:"interview_status" gorm:"type:varchar(255); default: unknown interview status"`
}

type Calendar struct {
	ID               uint   `json:"id" gorm:"primaryKey"`
	GoogleCalendarID string `json:"google_calendar_id" gorm:"not null"`
	Hosts            []Host `gorm:"many2many:host_calendars;joinForeignKey:host_id;inverseJoinForeignKey:calendar_id"`
}

type SampleMeetings struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Title       string `json:"title" gorm:"type:varchar(255); not null"`
	Description string `json:"description" gorm:"type:varchar(255); not null"`
	Duration    int    `json:"duration" gorm:"type:integer; not null"`
}

// Referential table connecting hosts to meetings. Hosts can have several meetings scheduled, and meetings can have several hosts.
type HostMeetings struct {
	HostID    uint `json:"host_id" gorm:"primaryKey; autoIncrement:false; not null"`
	MeetingID uint `json:"meeting_id" gorm:"primaryKey; autoIncrement:false; not null"`
}

// Referential table connecting hosts to time preferences. Hosts can have several time preferences.
type TimePreference struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	HostID      uint   `json:"host_id" gorm:"not null"` // Foreign key to hosts
	StartWindow string `json:"start_window" gorm:"type:string; default: 00:00"`
	EndWindow   string `json:"end_window" gorm:"type:string; default: 00:00"`
}

// Referential table connecting hosts to calendars. Hosts can have several individual calendars, and group calendars can have several hosts.
type HostCalendar struct {
	HostID     uint `json:"host_id" gorm:"primaryKey; autoIncrement:false; not null"`
	CalendarID uint `json:"calendar_id" gorm:"primaryKey; autoIncrement:false; not null"`
}
