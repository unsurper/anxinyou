package model

import "gorm.io/gorm"

type Filter struct {
	gorm.Model
	Word string `gorm:"not null"`
}
