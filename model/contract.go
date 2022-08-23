package model

import "gorm.io/gorm"

type Contract struct {
	gorm.Model
	Business  string  `gorm:"not null"` //导游
	Buyers    string  `gorm:"not null"` //用户
	Offer     float64 `gorm:"not null"` //劳务费
	Guarantee float64 `gorm:"not null"` //导游担保费
}
