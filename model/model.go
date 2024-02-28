package model

import "github.com/jinzhu/gorm"

type Meeting struct {
	gorm.Model
	Calendar    string `json:"calendar"`
	Duration    int    `json:"duration"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Hosts       string `json:"hosts"`
	HasBotGuest bool   `json:"hasBotGuest"`
}
