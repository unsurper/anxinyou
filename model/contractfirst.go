package model

import (
	"gorm.io/gorm"
	"time"
)

type Contractfirst struct {
	gorm.Model

	Business  string  `gorm:"not null"` //导游
	Buyers    string  `gorm:"not null"` //用户
	Offer     float64 `gorm:"not null"` //劳务费
	Guarantee float64 `gorm:"not null"` //导游担保费

	Contractimg string
	Starttime   time.Time `gorm:"type:datetime(3)"`
	Endtime     time.Time `gorm:"type:datetime(3)"`
	State       int       `gorm:"default:0;not null"`
}
