package model

import "gorm.io/gorm"

type Up struct {
	gorm.Model
	Upuser string `gorm:"not null"`
	Uppost string `gorm:"not null"`
}
