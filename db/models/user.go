package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	ChatID      int64  `json:"chat_id" gorm:"not null;column:chat_id"`
	Source      string `json:"source" gorm:"not null;column:source"`
	Destination string `json:"destination" gorm:"not null;column:destination"`
	Date        string `json:"date" gorm:"not null;column:date"`
}
