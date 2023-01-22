package controllers

import (
	"ayush-gupta-01/golang-api-tut/connect"
	"ayush-gupta-01/golang-api-tut/models"
	"ayush-gupta-01/golang-api-tut/utils"
	"fmt"
	"log"
	"net/http"
)

var gormDB = connect.GormDB

type response struct {
	Msg string `json:"msg"`
}

func SayHello(w http.ResponseWriter, r *http.Request) {
	res := response{"Hello World"}
	utils.WriteJsonData(w, res, 200)
}

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	if gormDB == nil {
		log.Fatal("Khali h ")
	}
	var res []models.Books
	result := gormDB.Find(&res)
	fmt.Println(result, res)
	utils.WriteJsonData(w, res, 200)
}

func GetBookById() {

}
