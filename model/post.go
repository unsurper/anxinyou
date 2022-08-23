package model

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Cover      string `gorm:"type:varchar(110)"`
	Telephone  string `gorm:"not null"`
	Text       string `gorm:"type:text"`
	Title      string `gorm:"not null"`
	Uppost     int    `gorm:"default:0;not null"` //点赞帖子
	Browsepost int    `gorm:"default:0;not null"` //浏览帖子
}
