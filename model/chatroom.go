package model

import "gorm.io/gorm"

type Chatroom struct {
	gorm.Model
	Newmsg string `gorm:"not null"`
	Who    string `gorm:"not null"`
	Usera  string `gorm:"not null"`
	Userb  string `gorm:"not null"`
}
