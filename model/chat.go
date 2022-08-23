package model

import "gorm.io/gorm"

type Chat struct {
	gorm.Model
	Sender   string `gorm:"not null"`
	Receiver string `gorm:"not null"`
	Content  string `gorm:"not null"`
}
