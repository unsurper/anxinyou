package model

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Telephone string `gorm:"not null"`
	Text      string `gorm:"type:text;not null"`
	Postid    string `gorm:"not null"`
	Upcomment int    `gorm:"default:0;not null"`
}
