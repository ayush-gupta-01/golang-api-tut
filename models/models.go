package models

import "github.com/jinzhu/gorm"

type Books struct {
	gorm.Model

	BookName     string `gorm:"column:name" json:"bookname"`
	Publications string `gorm:"column:publications" json:"publication"`
}
