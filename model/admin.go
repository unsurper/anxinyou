package model

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	Name     string `gorm:"type:varchar(20);not null;unique"`
	Password string `gorm:"size:255;not null"`
	Auth     int    `gorm:"default:0;not null"`
}
