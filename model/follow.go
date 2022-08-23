package model

import "gorm.io/gorm"

type Follow struct {
	gorm.Model
	Username string `gorm:"not null"`
	Follower string `gorm:"not null"`
}
