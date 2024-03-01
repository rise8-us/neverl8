package model

import "github.com/jinzhu/gorm"

type Meeting struct {
	gorm.Model
	Calendar    string
	Duration    int
	Title       string
	Description string
	Hosts       string
	HasBotGuest bool
}
