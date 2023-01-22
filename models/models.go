package models

type Books struct {
	Id           int64  `gorm:"column:id" json:"id"`
	BookName     string `gorm:"column:name" json:"bookname"`
	Publications string `gorm:"column:publications" json:"publications"`
}

type Response struct {
	ResponseMsg  string      `json:"responseMsg"`
	ResponseCode int         `json:"responseCode"`
	Data         interface{} `json:"data"`
	Error        interface{} `json:"error"`
}
