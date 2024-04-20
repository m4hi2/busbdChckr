package models

import "gorm.io/gorm"

type Log struct {
	gorm.Model

	ChatID int64  `json:"chat_id" gorm:"not null;column:chat_id"`
	Hash   string `json:"hash" gorm:"not null;column:hash;index"`
	Data   string `json:"data" gorm:"not null;column:data"`
}
