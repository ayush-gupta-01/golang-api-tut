package controllers

import (
	"ayush-gupta-01/golang-api-tut/connect"
	"ayush-gupta-01/golang-api-tut/models"
	"ayush-gupta-01/golang-api-tut/utils"
	"net/http"
)

var gormDB = connect.GormDB()

type response struct {
	Msg string `json:"msg"`
}

func SayHello(w http.ResponseWriter, r *http.Request) {
	res := response{"Hello World"}
	utils.WriteJsonData(w, res, 200)
}

func GetAllBooks(w http.ResponseWriter, r *http.Request) {

	var res []models.GetAllBooks
	result := gormDB.Find(&res)
	utils.WriteJsonData(w, result, 200)
}

func GetBookById() {

}
