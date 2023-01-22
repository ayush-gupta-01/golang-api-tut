package models

type Books struct {
	Id           uint64 `gorm:"column:id;primary_key" json:"id"`
	BookName     string `gorm:"column:name" json:"bookname"`
	Publications string `gorm:"column:publications" json:"publication"`
}
