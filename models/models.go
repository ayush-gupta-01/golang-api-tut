package models

type Books struct {
	Id           int64  `gorm:"column:id" json:"id"`
	BookName     string `gorm:"column:name" json:"bookname"`
	Publications string `gorm:"column:publications" json:"publications"`
}

type BookResp struct {
	Id          int64  `json:"id"`
	BookName    string `json:"bookname"`
	Publication string `json:"publication"`
	CreatedAt   string `json:"createdat"`
	UpdatedAt   string `json:"updatedat"`
	DeletedAt   string `json:"deletedat"`
}
