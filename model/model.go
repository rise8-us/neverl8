package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Meeting struct {
	gorm.Model
	Calendar    string `json:"calendar"`
	Duration    int    `json:"duration"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Hosts       string `json:"hosts"`
	HasBotGuest bool   `json:"has_bot_guest"`
}

type Meetings struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	CandidateId uint      `json:"candidate_id" gorm:"type:integer; not null"`
	Calendar    string    `json:"calendar" gorm:"type:varchar(255); not null"`
	Duration    int       `json:"duration" gorm:"type:integer; not null"`
	Title       string    `json:"title" gorm:"type:varchar(255); not null"`
	Description string    `json:"description" gorm:"type:varchar(255); default: no description provided"`
	HasBotGuest bool      `json:"has_bot_guest" gorm:"type:bool; default: false"`
	StartTime   time.Time `json:"start_time" gorm:"type: timestamp with time zone; not null"`
	EndTime     time.Time `json:"end_time" gorm:"type: timestamp with time zone; not null"`
	CreatedAt   time.Time `json:"created_at" gorm:"type: timestamp with time zone; not null; default: current_timestamp with time zone"`
	Hosts       []*Hosts  `json:"hosts" gorm:"many2many:host_meetings;"` // Many to many relationship with hosts
}

type Hosts struct {
	ID              uint              `json:"id" gorm:"primaryKey"`
	HostName        string            `json:"host_name" gorm:"type:varchar(255); not null"`
	Candidates      []Candidates      `json:"candidates_id" gorm:"foreignKey:HostId"`
	Meetings        []*Meetings       `json:"meetings" gorm:"many2many:host_meetings;"`  // Many to many relationship with meetings
	TimePreferences []TimePreferences `json:"time_preferences" gorm:"foreignKey:HostId"` // One to many relationship with time preferences
}

type Candidates struct {
	ID              uint   `json:"id" gorm:"primaryKey"`
	HostId          uint   `json:"host_id"` //foreign key to hosts
	CandidateName   string `json:"candidate_name" gorm:"type:varchar(255); not null"`
	Role            string `json:"role" gorm:"type:varchar(255); default: unknown role"`
	Email           string `json:"email" gorm:"type:varchar(255); default: unknown email"`
	PhoneNumber     string `json:"phone_number" gorm:"type:varchar(255); default: unknown phone number"`
	InterviewStatus string `json:"interview_status" gorm:"type:varchar(255); default: unknown interview status"`
}

// Referential table connecting hosts to meetings. Hosts can have several meetings scheduled, and meetings can have several hosts.
type HostMeetings struct {
	HostId    uint `json:"host_id" gorm:"primaryKey; autoIncrement:false; not null"`
	MeetingId uint `json:"meeting_id" gorm:"primaryKey; autoIncrement:false; not null"`
}

// Referential table connecting hosts to time preferences. Hosts can have several time preferences.
type TimePreferences struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	HostId      uint      `json:"host_id" gorm:"not null"` // Foreign key to hosts
	StartWindow time.Time `json:"start_window" gorm:"type: timestamp with time zone; not null"`
	EndWindow   time.Time `json:"end_window" gorm:"type: timestamp with time zone; not null"`
}
